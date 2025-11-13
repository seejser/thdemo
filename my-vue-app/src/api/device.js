import api from "@/api"; // 默认导入


export const getDeviceList = (params) => api.get("/device/list", { params });
export const getDeviceDetail = (id) => api.get(`/device/${id}`);
