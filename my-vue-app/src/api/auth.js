// src/api/auth.js
import api from "@/api"; // 默认导入


/**
 * 获取图片验证码
 * GET /api/v1/auth/captcha
 * @returns {Promise}
 */
export const getCaptcha = () => api.get("/auth/captcha");

/**
 * 发送邮箱验证码
 * GET /api/v1/auth/email_code
 * @param {Object} params - { email, captcha, captcha_id }
 */
export const sendEmailCode = (params) =>
  api.get("/auth/email_code", { params });

/**
 * 用户注册
 * POST /api/v1/auth/register
 * @param {Object} data - { username, password, email, verifyCode, captcha_id }
 */
export const register = (data) => api.post("/auth/register", data);

/**
 * 用户登录
 * POST /api/v1/auth/login
 * @param {Object} data - { username, password, captcha_id, captcha }
 */
export const login = (data) => api.post("/auth/login", data);

/**
 * 刷新 token
 * POST /api/v1/auth/refresh
 * @param {String} refreshToken
 */
export const refreshToken = (refreshToken) =>
  api.post("/auth/refresh", { refresh_token: refreshToken });
