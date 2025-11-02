import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Me from "../pages/Me.vue";
import Detail from "../pages/Detail.vue";
import Login from "../pages/Login.vue";
import Register from "../pages/Register.vue";
import NotFound from "../pages/NotFound.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: { requiresAuth: true, showTabBar: true },
  },
  {
    path: "/me",
    name: "Me",
    component: Me,
    meta: { requiresAuth: true, showTabBar: true },
  },
  {
    path: "/detail/:id",
    name: "Detail",
    component: Detail,
    meta: { requiresAuth: true, showTabBar: false },
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
    meta: { showTabBar: false },
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
    meta: { showTabBar: false },
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: NotFound,
    meta: { showTabBar: false },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const devMode = true; // 开发阶段直接放行
  const token = localStorage.getItem("token"); // 正式环境用真实登录状态

  if (devMode) return next(); // 开发模式直接跳转
  if (to.meta.requiresAuth && !token) return next("/login"); // 非开发模式才检查登录
  next();
});

export default router;
