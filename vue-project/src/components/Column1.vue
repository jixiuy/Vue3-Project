<template>
    <div class="column-1">
      <fabu-button></fabu-button>
    </div>

</template>

<script>
import FabuButton from './FabuButton.vue'
export default {
  components:{
    FabuButton
  }
}
</script>

<style scoped>

.column-1{
  flex: 20%;
  background-color: #f0f0f0; /* 设置背景色 */
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  padding-top: 40px;
}

</style>