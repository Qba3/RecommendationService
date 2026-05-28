<script setup>
import {
  cart,
  removeFromCart,
  increaseQuantity,
  decreaseQuantity,
  getTotal,
  clearCart
} from "../store/cart"

async function submitOrder() {
  const userId = Number(localStorage.getItem("userId"))

  const payload = {
    user_id: userId,
    items: cart.items.map(item => ({
      product_id: item.id,
      quantity: item.quantity
    }))
  }

  console.log(payload)

  const response = await fetch(
      "http://localhost:8080/order",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
      }
  )

  if (!response.ok) {
    const error = await response.text()

    console.error(error)

    alert("Błąd zapisu zamówienia")

    return
  }

  clearCart()

  alert("Zamówienie zapisane!")
}
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
          <h2 class="font-semibold">
            {{ item.name }}
          </h2>

          <div class="flex items-center gap-3 mt-2">
            <button
                @click="decreaseQuantity(item.id)"
                class="px-2 py-1 bg-gray-200 rounded"
            >
              -
            </button>

            <span>
              {{ item.quantity }}
            </span>

            <button
                @click="increaseQuantity(item.id)"
                class="px-2 py-1 bg-gray-200 rounded"
            >
              +
            </button>
          </div>
        </div>

        <div class="text-right">
          <p class="font-bold text-indigo-600">
            {{ item.price * item.quantity }} zł
          </p>

          <button
              @click="removeFromCart(item.id)"
              class="text-sm text-red-500"
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
            @click="submitOrder"
            class="mt-4 bg-green-600 text-white px-6 py-2 rounded-xl"
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