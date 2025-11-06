import { createWebHistory,createRouter } from "vue-router";

import Home from "./pages/Home.vue";
import My from "./pages/My.vue";
import Detail from "./pages/Detail.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/my", component: My },
  { path: "/detail", component: Detail },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
export default router;
