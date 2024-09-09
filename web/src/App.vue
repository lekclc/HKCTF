<script setup>
import { RouterLink, RouterView } from 'vue-router'
import LeftgoView from './views/LeftgoView.vue';
</script>


<template>
  <div id="LeftGo">
   <LeftgoView></LeftgoView>
  </div>
  <div id="Main">
    <RouterView />
  </div>
</template>

<style>
#LeftGo {
  position: fixed;
  top: 0;
  left: 0;
}
</style>

