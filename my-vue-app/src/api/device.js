import api from "./index.js";

export const getDeviceList = (params) => api.get("/device/list", { params });
export const getDeviceDetail = (id) => api.get(`/device/${id}`);
