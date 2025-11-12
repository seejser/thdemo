import api from "./index.js";

export const login = (data) => api.post("/auth/login", data);
export const register = (data) => api.post("/auth/register", data);
export const refreshToken = (refreshToken) =>
  api.post("/auth/refresh", { refresh_token: refreshToken });
