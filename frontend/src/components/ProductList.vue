<script setup>
import { ref, onMounted } from "vue"
import ProductCard from "./ProductCard.vue"

const products = ref([])
const emit = defineEmits(["select"])

onMounted(async () => {
  const res = await fetch("http://localhost:8080/products")
  products.value = await res.json()
})
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 text-indigo-700">
      Produkty
    </h2>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <ProductCard
          v-for="p in products"
          :key="p.id"
          :product="p"
          @select="emit('select', $event)"
      />
    </div>
  </div>
</template>