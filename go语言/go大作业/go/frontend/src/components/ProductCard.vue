    <template>
      <el-card :body-style="{ padding: '0px' }" class="product-card" shadow="hover" @click="goToDetail">
        <img :src="product.image" class="product-image" alt="商品图片" @error="e => e.target.src = '/default1.png'"/>
        <div style="padding: 14px">
          <span class="product-name">{{ product.name }}</span>
          <div class="bottom">
            <span class="product-price">¥{{ product.price }}</span>
            <el-button type="primary" text class="button" @click="$emit('buy-now', product)">立即购买</el-button>
            <el-button type="success" text class="button" @click="$emit('add-to-cart', product)">加入购物车</el-button>
          </div>
        </div>
      </el-card>
    </template>
    
    <script setup>
    import { useRouter } from 'vue-router';

    // 定义 props
    const props = defineProps({
      product: {
        type: Object,
        required: true,
      },
    });

    const router = useRouter();

    // 跳转到详情页
    const goToDetail = () => {
      router.push(`/product/${props.product.id}`);
    };
    </script>
    
    <style scoped>
    .product-card {
      width: 100%;
      cursor: pointer;
    }
    
    .product-image {
      width: 100%;
      height: 200px;
      object-fit: cover;
      display: block;
    }
    
    .product-name {
      font-size: 16px;
      color: #303133;
      display: block;
      margin-bottom: 10px;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
    
    .bottom {
      margin-top: 13px;
      line-height: 12px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      gap: 10px;
    }
    
    .product-price {
      font-size: 18px;
      color: #f56c6c;
      font-weight: bold;
    }
    
    .button {
      margin-left: 5px;
    }
    </style>