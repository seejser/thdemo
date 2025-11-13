// src/api/user.js
import api from "@/api"; // 默认导入


/**
 * 获取当前登录用户信息
 * GET /api/v1/user/info
 */
export const getUserInfo = () => api.get("/user/info");

/**
 * 下面是可扩展的用户接口示例
 * 如果将来后端增加修改用户信息、修改密码等接口，可以在这里封装
 */

// 更新用户信息（示例）
// export const updateUserInfo = (data) => api.put("/user/info", data);

// 修改密码（示例）
// export const changePassword = (data) => api.put("/user/password", data);
