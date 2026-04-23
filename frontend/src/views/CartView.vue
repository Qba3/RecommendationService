<script setup>
import {
  cart,
  removeFromCart,
  getTotal,
  increaseQuantity,
  decreaseQuantity
} from "../store/cart"
</script>

<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-2xl font-bold mb-6 text-indigo-700">
      Koszyk
    </h1>

    <div v-if="cart.items.length">
      <div
          v-for="item in cart.items"
          :key="item.id"
          class="bg-white p-4 rounded-xl shadow mb-4 flex justify-between items-center"
      >
        <div>
          <h2 class="font-semibold">{{ item.name }}</h2>

          <div class="flex items-center gap-3 mt-2">
            <button
                @click="decreaseQuantity(item.id)"
                class="px-2 py-1 bg-gray-200 rounded hover:bg-gray-300"
            >
              -
            </button>

            <span class="font-semibold">
              {{ item.quantity }}
            </span>

            <button
                @click="increaseQuantity(item.id)"
                class="px-2 py-1 bg-gray-200 rounded hover:bg-gray-300"
            >
              +
            </button>
          </div>
        </div>

        <div class="text-right">
          <p class="font-bold text-indigo-600">
            {{ item.id * item.quantity }} zł
          </p>

          <button
              @click="removeFromCart(item.id)"
              class="text-sm text-red-500 mt-1"
          >
            Usuń
          </button>
        </div>
      </div>

      <div class="text-right mt-6">
        <p class="text-xl font-bold">
          Suma: {{ getTotal() }} zł
        </p>

        <button
            class="mt-4 bg-green-600 text-white px-6 py-2 rounded-xl hover:bg-green-700"
        >
          Zamów
        </button>
      </div>
    </div>

    <div v-else class="text-gray-400">
      Koszyk jest pusty
    </div>
  </div>
</template>