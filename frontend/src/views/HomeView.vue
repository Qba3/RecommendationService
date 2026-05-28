<script setup>
import { onMounted, ref } from "vue"
import { addToCart } from "../store/cart"

const products = ref([])
const selectedProduct = ref(null)

async function loadRecommendations() {
  const userId = localStorage.getItem("userId")

  const response = await fetch(
      `http://localhost:8080/recommendations/${userId}`
  )

  products.value = await response.json()
}

function selectProduct(product) {
  selectedProduct.value = product
}

function handleAdd() {
  if (!selectedProduct.value) return

  addToCart(selectedProduct.value)
}

onMounted(() => {
  loadRecommendations()
})
</script>

<template>
  <div class="grid grid-cols-4 gap-6">

    <div class="col-span-3">
      <h1 class="text-3xl font-bold mb-6 text-indigo-700">
        Polecane produkty
      </h1>

      <div class="grid grid-cols-2 gap-6">
        <div
            v-for="product in products"
            :key="product.id"
            @click="selectProduct(product)"
            class="bg-white rounded-2xl shadow-lg p-6 border border-gray-100 cursor-pointer hover:border-indigo-400 transition"
        >
          <h2 class="text-xl font-semibold mb-2">
            {{ product.name }}
          </h2>

          <p class="text-indigo-600 font-bold text-lg">
            {{ product.price }} zł
          </p>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-2xl shadow-lg p-6 border border-gray-100">
      <h2 class="text-xl font-bold mb-4 text-indigo-600">
        Szczegóły produktu
      </h2>

      <div v-if="selectedProduct">
        <h3 class="text-lg font-semibold">
          {{ selectedProduct.name }}
        </h3>

        <p class="text-gray-500 mt-2">
          Produkt rekomendowany dla użytkownika.
        </p>

        <div class="mt-4 text-indigo-600 font-bold text-2xl">
          {{ selectedProduct.price }} zł
        </div>

        <button
            @click="handleAdd"
            class="mt-6 w-full bg-indigo-600 text-white py-2 rounded-xl hover:bg-indigo-700 transition"
        >
          Dodaj do koszyka
        </button>
      </div>

      <div v-else class="text-gray-400">
        Wybierz produkt z listy
      </div>
    </div>

  </div>
</template>