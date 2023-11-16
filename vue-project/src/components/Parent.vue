<template>
  <div class="container">
    <column1></column1>
    <column2></column2>
  </div>
</template>

<script>
import Column1 from "../components/Column1.vue";
import Column2 from "../components/Column2.vue";

export default {
  components:{
    Column1,
    Column2
  }
}
</script>

<style scoped>
.container{
  display: flex;
  height: 100%;
}
</style>