<template>
  <div
    class="w-full h-18 border border-[var(--el-border-color)] hover:border-[var(--el-text-color-disabled)] rounded-sm overflow-hidden"
  >
    <div
      v-if="!props.modelValue || props.modelValue.length === 0"
      class="flex items-center justify-center h-full cursor-pointer"
      @click.stop="openLibrary"
    >
      <el-icon :size="24" color="var(--el-text-color-placeholder)"
        ><Plus
      /></el-icon>
    </div>
    <div v-else class="h-full relative">
      <div
        class="flex absolute bg-[var(--el-color-danger)] right-0 top-0 z-10 rounded-bl-full cursor-pointer pl-0.5 pb-0.5"
      >
        <el-icon color="#ffffff" :size="16" @click="onImageRemove()"
          ><Close
        /></el-icon>
      </div>
      <el-image
        class="w-full h-full"
        preview-teleported
        :src="props.modelValue"
        :preview-src-list="[props.modelValue]"
        :zoom-rate="1.2"
        :max-scale="7"
        :min-scale="0.2"
        fit="cover"
      />
    </div>
  </div>
  <el-dialog
    v-model="isVisible"
    title="图片库"
    :width="860"
    align-center
    destroy-on-close
  >
    <component :is="imageLibrary" @confirm="onImageConfirm" />
  </el-dialog>
</template>

<script setup lang="ts">
import imageLibrary from "@/components/ImageLibrary.vue";
import { ref } from "vue";

const isVisible = ref(false);

const props = defineProps({
  modelValue: {
    type: String,
    default: "",
  },
  placeholder: {
    type: String,
    default: "",
  },
});

const emit = defineEmits(["update:modelValue"]);

const openLibrary = () => {
  isVisible.value = true;
};

const onImageRemove = () => {
  emit("update:modelValue", "");
};

const onImageConfirm = (url: string[]) => {
  emit("update:modelValue", url[0]);
  isVisible.value = false;
};

</script>

<style lang="scss" scoped></style>
