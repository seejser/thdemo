import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Me from "../pages/Me.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/me",
    name: "Me",
    component: Me,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
