<script setup>
import { ref } from "vue"
import { useRouter } from "vue-router"

const login = ref("")
const router = useRouter()

function handleLogin() {
  if (!login.value) return
  localStorage.setItem("userId", login.value)
  router.push("/home")
}
</script>

<template>
  <div class="flex items-center justify-center h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-md w-80">
      <h2 class="text-2xl font-bold mb-6 text-center">
        Recommender Shop
      </h2>

      <input
          v-model="login"
          placeholder="Podaj login"
          class="w-full border border-gray-300 rounded p-2 mb-4 focus:outline-none focus:ring-2 focus:ring-blue-400"
      />

      <button
          @click="handleLogin"
          class="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition"
      >
        Zaloguj
      </button>
    </div>
  </div>
</template>