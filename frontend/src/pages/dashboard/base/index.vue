<template>
  <div>
    <!-- 顶部 card  -->
    <top-panel class="row-container" />
    <!-- 中部图表  -->
    <middle-chart class="row-container" />
  </div>
</template>
<script setup lang="ts">
import MiddleChart from './components/MiddleChart.vue';
import TopPanel from './components/TopPanel.vue';

defineOptions({
  name: 'DashboardBase',
});
</script>
<style scoped>
.row-container:not(:last-child) {
  margin-bottom: 16px;
}
</style>
