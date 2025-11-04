import { createMemoryHistory, createRouter } from "vue-router";

import HomeView from "./pages/Home.vue";
import Detail from "./pages/Detail.vue";

const routes = [
  { path: "/", component: HomeView },
  { path: "/detail", component: Detail },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});
export default router;
