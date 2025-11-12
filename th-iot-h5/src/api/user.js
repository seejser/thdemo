import api from "./index.js";

export const getUserInfo = () => api.get("/user/info");
export const updateUserInfo = (data) => api.put("/user/info", data);
