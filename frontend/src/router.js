import { createRouter, createWebHistory } from "vue-router"
import LoginView from "./views/LoginView.vue"
import HomeView from "./views/HomeView.vue"
import CartView from "./views/CartView.vue"

const routes = [
    { path: "/", redirect: "/login" },
    { path: "/login", component: LoginView },
    { path: "/home", component: HomeView },
    { path: "/cart", component: CartView }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router