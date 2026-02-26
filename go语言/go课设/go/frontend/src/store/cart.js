import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import cartApi from '@/api/cart.js'

export const useCartStore = defineStore('cart', () => {
  // 购物车中的商品列表
  const items = ref([])

  // 计算购物车中商品总数量
  const totalItemsCount = computed(() => {
    return items.value.reduce((total, item) => total + item.quantity, 0)
  })

  // 计算购物车中商品总价
  const totalPrice = computed(() => {
    return items.value.reduce((total, item) => total + (item.product.price * item.quantity), 0)
  })

  // 从后端获取购物车
  async function fetchCart() {
    const res = await cartApi.getCart()
    if (res.code === 0) {
      items.value = res.data
    }
  }

  // 添加商品到购物车
  async function addItem(product, quantity) {
    const res = await cartApi.addToCart({ product_id: product.id, quantity })
    if (res.code === 0) {
      ElMessage.success({ message: `"${product.name}" 已加入购物车`, duration: 1500 })
      fetchCart()
    } else {
      ElMessage.error(res.msg || '加入购物车失败')
    }
  }

  // 从购物车移除商品
  async function removeItem(productId) {
    const res = await cartApi.deleteCartItem({ product_id: productId })
    if (res.code === 0) {
      fetchCart()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  }

  // 更新商品数量
  async function updateItemQuantity(productId, quantity) {
    const res = await cartApi.updateCart({ product_id: productId, quantity })
    if (res.code === 0) {
      fetchCart()
    } else {
      ElMessage.error(res.msg || '更新失败')
    }
  }

  // 清空购物车
  async function clearCart() {
    // 可遍历删除所有项
    for (const item of items.value) {
      await removeItem(item.product.id)
    }
    fetchCart()
  }

  return {
    items,
    totalItemsCount,
    totalPrice,
    addItem,
    removeItem,
    updateItemQuantity,
    clearCart,
    fetchCart,
  }
}) 