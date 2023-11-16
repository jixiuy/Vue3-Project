<template>
  <router-view :key="key" class="column-2"></router-view>
</template>

<script setup >
import {ref} from 'vue'
let key=ref(new Date())
</script>

<style scoped>
.column-2{
  flex: 80%;
  padding: 10px;
}
</style>