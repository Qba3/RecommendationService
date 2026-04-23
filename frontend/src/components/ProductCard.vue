<script setup>
const props = defineProps(["product"])
</script>

<template>
  <div
      class="bg-white rounded-2xl shadow-md hover:shadow-xl transition p-4 cursor-pointer border border-transparent hover:border-indigo-200"
      @click="$emit('select', product)"
  >
    <div class="h-40 bg-gradient-to-br from-indigo-100 to-blue-100 rounded-xl mb-3 flex items-center justify-center">
      <span class="text-indigo-300 font-semibold">IMG</span>
    </div>

    <h3 class="font-semibold text-lg">{{ product.name }}</h3>

    <p class="text-gray-500 text-sm mt-1">
      Super produkt
    </p>

    <div class="mt-3 text-indigo-600 font-bold">
      {{ product.id }} zł
    </div>
  </div>
</template>