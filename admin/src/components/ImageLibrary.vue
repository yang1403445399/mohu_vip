<template>
  <div class="flex">
    <div class="flex-1 border-r border-[var(--el-border-color)] flex flex-col">
      <div
        class="flex p-2 justify-center border-b border-[var(--el-border-color)]"
      >
        <div
          class="flex hover:text-[var(--el-color-primary)] cursor-pointer mx-2"
        >
          <el-tooltip content="新增目录">
            <el-button
              :icon="FolderAdd"
              size="small"
              type="success"
              plain
              @click="onColumnAdd()"
            />
          </el-tooltip>
        </div>
        <div
          class="flex hover:text-[var(--el-color-primary)] cursor-pointer mx-2"
        >
          <el-tooltip content="删除目录">
            <el-button :icon="FolderRemove" size="small" type="danger" plain @click="onColumnDelete()" />
          </el-tooltip>
        </div>
        <div
          class="flex hover:text-[var(--el-color-primary)] cursor-pointer mx-2"
        >
          <el-tooltip content="编辑目录">
            <el-button :icon="EditPen" size="small" type="primary" plain @click="onColumnEdit()" />
          </el-tooltip>
        </div>
      </div>
      <div class="flex-1 overflow-auto p-2">
        <div
          :class="[
            'flex items-center rounded-md px-4 hover:text-[var(--el-color-primary)] cursor-pointer',
            {
              'bg-[var(--el-color-primary-light-9)] text-[var(--el-color-primary)]':
                value.id === columnData.current,
            },
          ]"
          v-for="value in columnData.data"
          :key="value.id"
          @click="onColumnClick(value)"
        >
          <el-icon
            ><component
              :is="value.id === columnData.current ? FolderOpened : Folder"
          /></el-icon>
          <div
            class="text-sm leading-10 ml-1 overflow-hidden text-ellipsis whitespace-nowrap max-w-[100px]"
          >
            {{ value.name }}
          </div>
        </div>
      </div>
    </div>
    <div class="flex-4 flex flex-col h-119">
      <div
        class="flex justify-between items-center p-2 border-b border-[var(--el-border-color)]"
      >
        <div class="flex items-center">
          <el-checkbox
            v-model="imageData.checked.all"
            @change="onImageCheckedChange"
            size="small"
            >全选</el-checkbox
          >
        </div>
        <div class="flex">
          <el-button size="small" type="success" plain @click="onImageAdd()"
            >新增</el-button
          >
          <el-button
            size="small"
            type="danger"
            plain
            @click="onImageDelete(imageData.checked.list)"
            >删除</el-button
          >
          <el-button
            size="small"
            type="primary"
            plain
            @click.stop.prevent="onImageBatchUse(imageData.checked.list)"
            >使用</el-button
          >
        </div>
      </div>
      <div class="flex-1 overflow-auto" v-loading="imageData.load">
        <el-checkbox-group
          v-model="imageData.checked.list"
          size="small"
          class="grid grid-cols-4 gap-2 p-2"
          v-if="imageData.data.list.length > 0"
        >
          <el-checkbox
            v-for="value in imageData.data.list"
            :key="value.id"
            :value="value.id"
            border
          >
            <el-popover :width="200" :hide-after="0" trigger="hover">
              <template #reference>
                <el-image class="w-full h-full" :src="value.url" fit="cover" />
              </template>
              <div class="flex justify-center">
                <el-button
                  size="small"
                  type="success"
                  plain
                  :disabled="imageData.disable"
                  @click.stop.prevent="onImagePreview(value.url)"
                  >预览</el-button
                >
                <el-button
                  size="small"
                  type="danger"
                  plain
                  :disabled="imageData.disable"
                  @click.stop.prevent="onImageDelete(value.id)"
                  >删除</el-button
                >
                <el-button
                  size="small"
                  type="primary"
                  plain
                  :disabled="imageData.disable"
                  @click.stop.prevent="onImageUse(value)"
                  >使用</el-button
                >
              </div>
            </el-popover>
          </el-checkbox>
        </el-checkbox-group>
        <div class="h-full flex items-center justify-center" v-else>
          <el-empty description="暂无图片" :image-size="100" />
        </div>
      </div>
      <div
        class="flex justify-end py-3 px-2 border-t border-[var(--el-border-color)]"
      >
        <el-pagination
          small
          background
          layout="total, sizes, prev, pager, next, jumper"
          v-model:current-page="imageData.data.current"
          v-model:page-size="imageData.data.size"
          :page-sizes="[12, 24, 48, 96]"
          :total="imageData.data.total"
          @current-change="onImageCurrentChange"
          @size-change="onImageSizeChange"
        />
      </div>
    </div>
  </div>
  <el-dialog
    v-model="columnData.dialog.show"
    align-center
    :title="columnData.dialog.title"
    :width="500"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRule"
      class="p-4"
    >
      <el-form-item prop="name">
        <el-input v-model="formData.name" placeholder="请输入栏目名称" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="">取消</el-button>
      <el-button type="primary" @click="onColumnSubmit(formRef)"
        >确定</el-button
      >
    </template>
  </el-dialog>
  <el-dialog
    v-model="imageData.dialog.show"
    align-center
    title="新增图片"
    :width="500"
  >
    <div class="p-4">
      <el-upload
        ref="uploadRef"
        drag
        multiple
        :auto-upload="false"
        :show-file-list="false"
        :limit="8"
        :before-upload="onImageBeforeUpload"
        :on-change="onImageUpload"
        :on-exceed="onImageExceed"
        accept=".jpg,.jpeg,.png,.gif,.webp"
      >
        <el-icon :size="60"><UploadFilled /></el-icon>
        <div class="text-sm">
          将文件拖放到此处或<em class="text-[var(--el-color-primary)]"
            >单击上传</em
          >
        </div>
        <template #tip>
          <div class="text-xs mt-2">jpg/png/gif/webp文件大小不超过2MB</div>
        </template>
      </el-upload>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ElMessage, ElMessageBox } from "element-plus";
import { onMounted, reactive, ref } from "vue";
import {
  EditPen,
  Folder,
  FolderOpened,
  FolderAdd,
  FolderRemove,
} from "@element-plus/icons-vue";
import {
  reqImageList,
  reqImageUpload,
  reqImageDelete,
  reqImageColumnList,
  reqImageColumnSubmit,
  reqImageColumnDelete,
} from "@/api";
import type {
  ResponseData,
  PaginationData,
  ImageData,
  ImageColumnData,
} from "@/types";
import type {
  UploadFile,
  UploadRawFile,
  UploadInstance,
  FormInstance,
  FormRules,
} from "element-plus";

const props = defineProps({
  multiple: {
    type: Boolean,
    default: false,
  },
});

const emits = defineEmits(["confirm"]);

const uploadRef = ref<UploadInstance>();
const formRef = ref<FormInstance>();
const formData = reactive<ImageColumnData>({
  name: "",
});
const formRule = reactive<FormRules<ImageColumnData>>({
  name: [{ required: true, message: "请输入栏目名称", trigger: "blur" }],
});
const columnData = reactive({
  load: true,
  data: [] as ImageColumnData[],
  current: 0,
  dialog: {
    title: "新增目录",
    show: false,
  },
});
const imageData = reactive({
  load: true,
  data: {
    current: 1,
    total: 0,
    size: 12,
    list: [],
  } as PaginationData<ImageData[]>,
  disable: false,
  checked: {
    all: false,
    list: [] as number[],
  },
  dialog: {
    show: false,
  },
});

const getImageColumnList = async () => {
  const response: ResponseData<ImageColumnData[]> = await reqImageColumnList();
  if (response.code === 200) {
    columnData.data = response.data!;
    if (columnData.data.length > 0) {
      columnData.current = columnData.data[0].id!;
    }
  }
  columnData.load = false;
};

const onColumnClick = (item: ImageColumnData) => {
  columnData.current = item.id!;
  imageData.load = true;
  imageData.checked.list = [];
  imageData.data.current = 1;
  getImageList();
};

const onColumnAdd = () => {
  columnData.dialog.show = true;
};

const onColumnDelete = () => {
  if (columnData.current === 0) {
    ElMessage.error("没有目录可删除");
    return;
  }

  if (imageData.data.total > 0) {
    ElMessage.error("目录下有图片，无法删除");
    return;
  }

  ElMessageBox.alert("是否确认删除？删除后将无法恢复！", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async() => {
    try {
      const response: ResponseData = await reqImageColumnDelete([columnData.current]);
      if (response.code === 200) {
        ElMessage({
          message: response.msg,
          type: "success",
        });
        columnData.load = true;
        imageData.load = true;
        await getImageColumnList();
        await getImageList();
      }
    } catch (error: any) {
      ElMessage({
        message: error.message,
        type: "error",
      });
    }
  })
}

const onColumnSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;

  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const response: ResponseData = await reqImageColumnSubmit(
          formData
        );
        if (response.code === 200) {
          ElMessage({
            message: response.msg,
            type: "success",
          });
          formRef.value?.resetFields();
          columnData.dialog.show = false;
          columnData.load = true;
          await getImageColumnList();
        }
      } catch (error: any) {
        ElMessage({
          message: error.message,
          type: "error",
        });
      }
    }
  });
};

const getImageList = async () => {
  const response: ResponseData<PaginationData<ImageData[]>> =
    await reqImageList({
      current: imageData.data.current,
      size: imageData.data.size,
      column_id: columnData.current,
    });
  if (response.code === 200) {
    imageData.data = response.data!;
  }
  imageData.load = false;
};

const onImageCurrentChange = async () => {
  imageData.load = true;
  imageData.checked.list = [];
  await getImageList();
};

const onImageSizeChange = async () => {
  imageData.load = true;
  imageData.checked.list = [];
  imageData.data.current = 1;
  await getImageList();
}

const onImageCheckedChange = () => {
  imageData.checked.list = imageData.checked.all
    ? imageData.data.list.map((item) => item.id)
    : [];
};

const onImageAdd = () => {
  if (columnData.current === 0) {
    ElMessage.error("请先创建目录");
    return;
  }
  imageData.dialog.show = true;
};

const onColumnEdit = () => {
  if (columnData.current === 0) {
    ElMessage.error("没有目录可编辑");
    return;
  }
  formData.id = columnData.current;
  formData.name = columnData.data.find((item) => item.id === columnData.current)!.name;
  columnData.dialog.show = true;
};

const onImageDelete = async (id: number[] | number) => {
  if (!Array.isArray(id)) {
    id = [id];
  }
  if (id.length === 0) {
    ElMessage.error("请选择图片");
    return;
  }

  ElMessageBox.alert("是否确认删除？删除后将无法恢复！", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    try {
      const response: ResponseData = await reqImageDelete(id);
      if (response.code === 200) {
        imageData.checked.list = [];
        imageData.load = true;
        ElMessage({
          message: response.msg,
          type: "success",
          onClose: async () => {
            imageData.checked.all = false;
            await getImageList();
          },
        });
      }
    } catch (error: any) {
      ElMessage({
        message: error.message,
        type: "error",
      });
    }
  });
};

const onImageBeforeUpload = (file: UploadRawFile) => {
  const imageTypes = ["image/jpeg", "image/png", "image/gif", "image/webp"];
  const isAllowedType = imageTypes.includes(file.type);
  const isAllowedSize = file.size / 1024 / 1024 < 2;

  if (!isAllowedType) {
    ElMessage.error("上传图片只能是JPG、PNG、GIF、WebP格式");
  }

  if (!isAllowedSize) {
    ElMessage.error("上传图片大小不能超过2MB");
  }

  return isAllowedType && isAllowedSize;
};

const onImageUpload = async (file: UploadFile) => {
  try {
    const response: ResponseData = await reqImageUpload({
      image: file.raw!,
      column_id: columnData.current,
    });
    if (response.code === 200) {
      uploadRef.value!.clearFiles();
      imageData.dialog.show = false;
      imageData.load = true;
      imageData.checked.list = [];
      ElMessage({
        message: response.msg,
        type: "success",
        onClose: async () => {
          await getImageList();
        },
      });
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const onImageExceed = () => {
  ElMessage.error("超出单次上传数量");
};

const onImagePreview = (url: string) => {
  window.open(url, "_blank");
};

const onImageUse = (data: ImageData) => {
  emits("confirm", [data.url]);
};

const onImageBatchUse = (data: number[]) => {
  if (!props.multiple && data.length > 1) {
    ElMessage.error("请选择单张图片");
    return;
  }

  if (data.length === 0) {
    ElMessage.error("请选择图片");
    return;
  }

  emits(
    "confirm",
    imageData.data.list
      .filter((item) => imageData.checked.list.includes(item.id))
      .map((item) => item.url)
  );
};

const loadPageData = async () => {
  try {
    await getImageColumnList();
    await getImageList();
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

onMounted(async () => {
  await loadPageData();
});
</script>

<style lang="scss" scoped>
:deep(.el-checkbox-group) {
  .el-checkbox {
    position: relative;
    height: auto;
    margin: 0;
    background: var(--grey-bg-color);

    &.is-bordered {
      padding: 4px;
    }

    .el-checkbox__input {
      position: absolute;
      top: 8px;
      left: 8px;
    }

    .el-checkbox__label {
      width: 100%;
      height: 108px;
      padding: 0;

      .el-dropdown {
        width: 100%;
        height: 100%;
      }
    }
  }
}
</style>
