<template>
  <el-input
    class="cursor-pointer"
    :value="props.modelValue"
    :placeholder="placeholder"
    readonly
    @focus="onImagePreview()"
  >
    <template #append>
      <el-button :icon="Pointer" @click.stop="openLibrary" />
    </template>
  </el-input>
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
import { Pointer } from "@element-plus/icons-vue";

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

const onImagePreview = () => {
  if (props.modelValue) {
    window.open(props.modelValue);
  }
}

const onImageConfirm = (url: string[]) => {
  emit("update:modelValue", url[0]);
  isVisible.value = false;
};
</script>

<style lang="scss" scoped></style>
