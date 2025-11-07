import { createWebHistory,createRouter } from "vue-router";

import Home from "./pages/Home.vue";
import My from "./pages/My.vue";
import Detail from "./pages/Detail.vue";
import Login from "./pages/Login.vue";
import Register from "./pages/Register.vue";
import NotFound from "./pages/404.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/my", component: My },
  { path: "/login", component: Login },
  { path: "/register", component: Register },
  { path: "/404", component: NotFound },
  { path: "/detail/:id", component: Detail },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
export default router;
