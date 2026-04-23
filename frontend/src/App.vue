<script setup>
import { cart } from "./store/cart"
</script>

<template>
  <div>
    <header class="bg-white shadow px-6 py-4 flex justify-between items-center">
      <h1 class="text-xl font-bold text-indigo-600">
        Recommender Shop
      </h1>

      <div class="flex items-center gap-4">
        <router-link to="/cart" class="relative">
          🛒
          <span
              v-if="cart.items.length"
              class="absolute -top-2 -right-2 bg-indigo-600 text-white text-xs px-2 rounded-full"
          >
            {{ cart.items.length }}
          </span>
        </router-link>

        <button class="text-sm bg-gray-100 px-3 py-1 rounded hover:bg-gray-200">
          Wyloguj
        </button>
      </div>
    </header>

    <main class="p-6">
      <router-view />
    </main>
  </div>
</template>