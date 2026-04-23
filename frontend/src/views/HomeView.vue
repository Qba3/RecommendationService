<script setup>
import { ref } from "vue"
import ProductList from "../components/ProductList.vue"
import { addToCart } from "../store/cart"

const selectedProduct = ref(null)

function selectProduct(product) {
  selectedProduct.value = product
}

function handleAdd() {
  if (selectedProduct.value) {
    addToCart(selectedProduct.value)
  }
}
</script>

<template>
  <div class="grid grid-cols-4 gap-6">

    <div class="col-span-3">
      <ProductList @select="selectProduct" />
    </div>

    <div class="bg-white rounded-2xl shadow-lg p-6 border border-gray-100">
      <h2 class="text-xl font-bold mb-4 text-indigo-600">
        Szczegóły
      </h2>

      <div v-if="selectedProduct">
        <h3 class="text-lg font-semibold">
          {{ selectedProduct.name }}
        </h3>

        <p class="text-gray-500 mt-2">
          Opis produktu
        </p>

        <div class="mt-4 text-indigo-600 font-bold text-xl">
          {{ selectedProduct.id }} zł
        </div>

        <button
            @click="handleAdd"
            class="mt-6 w-full bg-indigo-600 text-white py-2 rounded-xl hover:bg-indigo-700 transition"
        >
          Dodaj do koszyka
        </button>
      </div>

      <div v-else class="text-gray-400">
        Wybierz produkt
      </div>
    </div>

  </div>
</template>