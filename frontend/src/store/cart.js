import { reactive } from "vue"

export const cart = reactive({
    items: []
})

export function addToCart(product) {
    const existing = cart.items.find(p => p.id === product.id)

    if (existing) {
        existing.quantity += 1
    } else {
        cart.items.push({ ...product, quantity: 1 })
    }
}

export function increaseQuantity(id) {
    const item = cart.items.find(p => p.id === id)
    if (item) item.quantity += 1
}

export function decreaseQuantity(id) {
    const item = cart.items.find(p => p.id === id)
    if (!item) return

    if (item.quantity > 1) {
        item.quantity -= 1
    } else {
        cart.items = cart.items.filter(p => p.id !== id)
    }
}

export function removeFromCart(id) {
    cart.items = cart.items.filter(p => p.id !== id)
}

export function getTotal() {
    return cart.items.reduce((sum, p) => sum + p.id * p.quantity, 0)
}