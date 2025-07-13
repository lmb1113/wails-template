<template>
  <div v-if="showFrame" class="content">
    <template v-for="frame in getFramePages" :key="frame.path">
      <frame-content v-if="hasRenderFrame(frame.name)" v-show="showIframe(frame)" :frame-src="frame.meta.frameSrc" />
    </template>
  </div>
</template>
<script lang="ts">
import { computed, defineComponent, unref } from 'vue';

import FrameContent from '../components/FrameContent.vue';
import { useFrameKeepAlive } from './useFrameKeepAlive';

export default defineComponent({
  name: 'FrameLayout',
  components: { FrameContent },
  setup() {
    const { getFramePages, hasRenderFrame, showIframe } = useFrameKeepAlive();

    const showFrame = computed(() => unref(getFramePages).length > 0);

    return { getFramePages, hasRenderFrame, showIframe, showFrame };
  },
});
</script>
<style scoped>
.content {
  overflow-y: scroll;
  scrollbar-width: none;
  //margin: -20px;
  margin: calc(0px - var(--td-comp-paddingTB-xl)) calc(0px - var(--td-comp-paddingLR-xl));
}
</style>
