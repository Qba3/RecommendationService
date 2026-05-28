<script setup>
import { ref } from "vue"
import { useRouter } from "vue-router"

const login = ref("")
const router = useRouter()

async function handleLogin() {
  if (!login.value) return

  const response = await fetch("http://localhost:8080/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      username: login.value
    })
  })

  if (!response.ok) {
    alert("Błąd logowania")
    return
  }

  const user = await response.json()
  console.log(user)
  localStorage.setItem("userId", user.id)
  localStorage.setItem("username", user.username)

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
          class="w-full border border-gray-300 rounded p-2 mb-4"
      />

      <button
          @click="handleLogin"
          class="w-full bg-blue-500 text-white py-2 rounded"
      >
        Zaloguj
      </button>
    </div>
  </div>
</template>