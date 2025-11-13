import axios from "axios";
import { showNotify } from "@nutui/nutui";

const baseURL = import.meta.env.VITE_APP_API_URL
  ? `${import.meta.env.VITE_APP_API_URL}/api/v1`
  : "/api/v1";

const api = axios.create({
  baseURL,
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});


function handleLoginExpired() {
  localStorage.removeItem("token");
  localStorage.removeItem("refreshToken");
  showNotify.success("登录状态已过期，请重新登录");
  window.location.href = "/login.html";
}

api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token) config.headers["Authorization"] = "Bearer " + token;
    return config;
  },
  (error) => Promise.reject(error)
);

let isRefreshing = false;
let refreshSubscribers = [];
function subscribeTokenRefresh(cb) { refreshSubscribers.push(cb); }
function onRefreshed(token) { refreshSubscribers.forEach(cb => cb(token)); refreshSubscribers = []; }

api.interceptors.response.use(
  (response) => {
    if (response.data && response.data.code !== 0) {
      const msg = response.data.message || response.data.msg || "未知业务错误";
      showNotify.success(msg);
      const err = new Error(msg);
      err.code = response.data.code;
      return Promise.reject(err);
    }
    return response;
  },
  async (error) => {
    const originalRequest = error.config;
    if (!error.response || error.response.status !== 401) return Promise.reject(error);

    const refreshToken = localStorage.getItem("refreshToken");
    if (!refreshToken) { handleLoginExpired(); return Promise.reject(error); }

    if (!isRefreshing) {
      isRefreshing = true;
      try {
        const res = await axios.post("/api/v1/auth/refresh", { refresh_token: refreshToken });
        if (res.data.code === 0 && res.data.data?.accessToken) {
          const newToken = res.data.data.accessToken;
          const newRefreshToken = res.data.data.refreshToken || refreshToken;
          localStorage.setItem("token", newToken);
          localStorage.setItem("refreshToken", newRefreshToken);
          isRefreshing = false;
          onRefreshed(newToken);
        } else { handleLoginExpired(); }
      } catch (e) { handleLoginExpired(); } finally { isRefreshing = false; }
    }

    return new Promise(resolve => {
      subscribeTokenRefresh((token) => {
        originalRequest.headers["Authorization"] = "Bearer " + token;
        resolve(api.request(originalRequest));
      });
    });
  }
);

window.api = api;
export default api;


