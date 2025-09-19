<script setup lang="ts">
import { Editor, EditorContent } from "@tiptap/vue-3";
import {
  FontFamily,
  TextStyle,
  FontSize,
  Color,
  BackgroundColor,
} from "@tiptap/extension-text-style";
import TextAlign from "@tiptap/extension-text-align";
import StarterKit from "@tiptap/starter-kit";
import { ref, reactive, onBeforeUnmount, onMounted } from "vue";
import type { Level } from "@tiptap/extension-heading";

const props = defineProps({
  modelValue: {
    type: String,
    default: "",
  },
});

const emits = defineEmits(["update:modelValue"]);

const fontFamilyData = ref([
  { label: "Arial", value: "Arial, sans-serif" },
  { label: "Times New Roman", value: "Times New Roman, serif" },
  { label: "Helvetica", value: "Helvetica, sans-serif" },
  { label: "Georgia", value: "Georgia, serif" },
  { label: "Verdana", value: "Verdana, sans-serif" },
  { label: "Courier New", value: "Courier New, monospace" },
  { label: "微软雅黑", value: '"Microsoft YaHei", sans-serif' },
  { label: "宋体", value: "SimSun, serif" },
  { label: "黑体", value: "SimHei, sans-serif" },
  { label: "楷体", value: "KaiTi, serif" },
  { label: "仿宋", value: "FangSong, serif" },
]);

const fontSizeData = ref([
  { label: "一号", value: "34px" },
  { label: "小一", value: "32px" },
  { label: "二号", value: "30px" },
  { label: "小二", value: "28px" },
  { label: "三号", value: "26px" },
  { label: "小三", value: "24px" },
  { label: "四号", value: "22px" },
  { label: "小四", value: "20px" },
  { label: "五号", value: "18px" },
  { label: "小五", value: "16px" },
  { label: "六号", value: "14px" },
  { label: "小六", value: "12px" },
]);

const fontColorData = reactive({
  popover: false,
  current: "#000000",
  list: [
    "#f6f6f6",
    "#000000",
    "#ff4500",
    "#ff8c00",
    "#ffd700",
    "#90ee90",
    "#00ced1",
    "#1e90ff",
    "#c71585",
  ],
});

const fontBackgroundData = reactive({
  popover: false,
  current: "#000000",
  list: [
    "#f6f6f6",
    "#000000",
    "#ff4500",
    "#ff8c00",
    "#ffd700",
    "#90ee90",
    "#00ced1",
    "#1e90ff",
    "#c71585",
  ],
});

const fontHeadingData = ref([
  { label: "一级标题", value: 1 },
  { label: "二级标题", value: 2 },
  { label: "三级标题", value: 3 },
  { label: "四级标题", value: 4 },
  { label: "五级标题", value: 5 },
  { label: "六级标题", value: 6 },
]);

const onFontFamilyApply = (family: string) => {
  editor.commands.setFontFamily(family);
};

const onFontSizeApply = (size: string) => {
  editor.commands.setFontSize(size);
};

const onFontColorPopoverSwitch = (event: MouseEvent) => {
  event.stopPropagation();

  const button = document.querySelector("#font-color-button");
  const apply = document.querySelector("#font-color-apply");
  const icon = document.querySelector("#font-color-icon");

  if (button === event.target || icon === event.target) {
    fontColorData.popover = !fontColorData.popover;
  }

  if (!button?.contains(event.target as Node)) {
    fontColorData.popover = false;
  }

  if (apply?.contains(event.target as Node)) {
    editor.commands.setColor(fontColorData.current);
    fontColorData.popover = false;
  }
};

const onFontBackgroundPopoverSwitch = (event: MouseEvent) => {
  event.stopPropagation();

  const button = document.querySelector("#font-background-button");
  const apply = document.querySelector("#font-background-apply");
  const icon = document.querySelector("#font-background-icon");

  if (button === event.target || icon === event.target) {
    fontBackgroundData.popover = !fontBackgroundData.popover;
  }

  if (!button?.contains(event.target as Node)) {
    fontBackgroundData.popover = false;
  }

  if (apply?.contains(event.target as Node)) {
    editor.commands.setBackgroundColor(fontBackgroundData.current);
    fontBackgroundData.popover = false;
  }
};

const onFontHeadingApply = (heading: Level) => {
  editor.commands.setHeading({ level: heading });
};

const onFontBoldApply = () => {
  editor.commands.toggleBold();
};

const onFontItalicApply = () => {
  editor.commands.toggleItalic();
};

const onFontUnderlineApply = () => {
  editor.commands.toggleUnderline();
};

const onFontStrikethroughApply = () => {
  editor.commands.toggleStrike();
};

const onFontAlignApply = (align: string) => {
  console.log(align);
  editor.commands.setTextAlign(align);
};

const onListBulletApply = () => {
  editor.commands.toggleBulletList();
};

const onListOrderApply = () => {
  editor.commands.toggleOrderedList();
};

const editor = new Editor({
  content: props.modelValue || "<p>I'm running Tiptap with Vue.js. 🎉</p>",
  extensions: [
    StarterKit,
    FontFamily,
    TextStyle,
    FontSize,
    Color,
    BackgroundColor,
    TextAlign.configure({
      types: ["heading", "paragraph"],
    }),
  ],
});

onMounted(() => {
  document.addEventListener("click", onFontColorPopoverSwitch);
  document.addEventListener("click", onFontBackgroundPopoverSwitch);
});

onBeforeUnmount(() => {
  editor.destroy();
  document.removeEventListener("click", onFontColorPopoverSwitch);
  document.removeEventListener("click", onFontBackgroundPopoverSwitch);
});
</script>

<template>
  <div class="w-full border border-[var(--el-border-color)] rounded-sm relative">
    <div class="flex flex-wrap items-start py-2 px-1 bg-[var(--grey-bg-color)]">
      <div class="flex cursor-pointer">
        <el-tooltip content="源代码">
          <el-button size="small" text
            ><img class="w-4 h-4" src="@/assets/icons/html.svg" alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="大小">
          <el-dropdown trigger="click" @command="onFontSizeApply">
            <el-button size="small" text
              ><img class="w-4 h-4" src="@/assets/icons/font_size.svg" alt=""
            /></el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="value in fontSizeData"
                  :key="value.value"
                  :command="value.value"
                  :style="{ fontSize: value.value }"
                  >{{ value.label }}</el-dropdown-item
                >
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="字体">
          <el-dropdown trigger="click" @command="onFontFamilyApply">
            <el-button size="small" text
              ><img class="w-5 h-5" src="@/assets/icons/font_family.svg" alt=""
            /></el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="value in fontFamilyData"
                  :key="value.value"
                  :command="value.value"
                  :style="{ fontFamily: value.value }"
                  >{{ value.label }}</el-dropdown-item
                >
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="颜色">
          <el-button id="font-color-button" size="small" text>
            <el-popover
              :teleported="false"
              :visible="fontColorData.popover"
              :width="352"
            >
              <template #reference>
                <img
                  id="font-color-icon"
                  class="w-5 h-5"
                  src="@/assets/icons/font_color.svg"
                  alt=""
                />
              </template>
              <template #default>
                <div class="w-full relative">
                  <el-color-picker-panel
                    v-model="fontColorData.current"
                    show-alpha
                    :predefine="fontColorData.list"
                  />
                  <el-button
                    id="font-color-apply"
                    class="absolute bottom-3 right-3"
                    size="small"
                  >
                    确定
                  </el-button>
                </div>
              </template>
            </el-popover>
          </el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="背景">
          <el-button id="font-background-button" size="small" text>
            <el-popover
              :teleported="false"
              :visible="fontBackgroundData.popover"
              :width="352"
            >
              <template #reference>
                <img
                  id="font-background-icon"
                  class="w-5 h-5"
                  src="@/assets/icons/font_background.svg"
                  alt=""
                />
              </template>
              <template #default>
                <div class="w-full relative">
                  <el-color-picker-panel
                    v-model="fontBackgroundData.current"
                    show-alpha
                    :predefine="fontBackgroundData.list"
                  />
                  <el-button
                    id="font-background-apply"
                    class="absolute bottom-3 right-3"
                    size="small"
                  >
                    确定
                  </el-button>
                </div>
              </template>
            </el-popover>
          </el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="标题">
          <el-dropdown trigger="click" @command="onFontHeadingApply">
            <el-button size="small" text
              ><img
                class="w-5 h-5"
                src="@/assets/icons/font_heading.svg"
                alt=""
            /></el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="(value, index) in fontHeadingData"
                  :key="value.value"
                  :command="value.value"
                  class="font-bold"
                  :style="{ fontSize: (9 - index) * 4 + `px` }"
                  >{{ value.label }}</el-dropdown-item
                >
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="加粗">
          <el-button size="small" text @click="onFontBoldApply()"
            ><img class="w-4.5 h-4.5" src="@/assets/icons/font_bold.svg" alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="斜体">
          <el-button size="small" text @click="onFontItalicApply()"
            ><img class="w-5 h-5" src="@/assets/icons/font_italic.svg" alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="下划线">
          <el-button size="small" text @click="onFontUnderlineApply()"
            ><img
              class="w-5 h-5"
              src="@/assets/icons/font_underline.svg"
              alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="删除线">
          <el-button size="small" text @click="onFontStrikethroughApply()"
            ><img
              class="w-5 h-5"
              src="@/assets/icons/font_strikethrough.svg"
              alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="对齐">
          <el-dropdown trigger="click" @command="onFontAlignApply">
            <el-button size="small" text
              ><img class="w-5 h-5" src="@/assets/icons/text_align.svg" alt=""
            /></el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="left">左对齐</el-dropdown-item>
                <el-dropdown-item command="center">居中对齐</el-dropdown-item>
                <el-dropdown-item command="right">右对齐</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="有序列表">
          <el-button size="small" text @click="onListOrderApply()"
            ><img class="w-5 h-5" src="@/assets/icons/list_order.svg" alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="无序列表">
          <el-button size="small" text @click="onListBulletApply()"
            ><img class="w-5 h-5" src="@/assets/icons/list_bullet.svg" alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <div class="flex cursor-pointer">
        <el-tooltip content="超链接">
          <el-button size="small" text @click="onListOrderApply()"
            ><img class="w-5 h-5" src="@/assets/icons/link.svg" alt=""
          /></el-button>
        </el-tooltip>
      </div>
      <!-- <div class="flex m-2 cursor-pointer">
        <el-tooltip content="代码">
          <img class="w-5 h-5" src="@/assets/icons/code.svg" alt="" />
        </el-tooltip>
      </div>
      <div class="flex m-2 cursor-pointer">
        <el-tooltip content="表格">
          <img class="w-5 h-5" src="@/assets/icons/table.svg" alt="" />
        </el-tooltip>
      </div>
      <div class="flex m-2 cursor-pointer">
        <el-tooltip content="图片">
          <img class="w-5 h-5" src="@/assets/icons/image.svg" alt="" />
        </el-tooltip>
      </div>
      <div class="flex m-2 cursor-pointer">
        <el-tooltip content="视频">
          <img class="w-5 h-5" src="@/assets/icons/video.svg" alt="" />
        </el-tooltip>
      </div>
      <div class="flex m-2 cursor-pointer">
        <el-tooltip content="引用">
          <img class="w-5 h-5" src="@/assets/icons/quote.svg" alt="" />
        </el-tooltip>
      </div> -->
    </div>
    <div class="border-t border-[var(--el-border-color)] px-3 h-146">
      <editor-content :editor="editor" />
    </div>
    <div class="border-t border-[var(--el-border-color)] bg-[var(--grey-bg-color)] leading-none h-1 flex items-center justify-center">
      <div class="w-4 h-[1px] bg-[var(--el-border-color)]"></div>
    </div>
    <div class="absolute bottom-2 right-1 leading-none text-xs text-[var(--el-color-info)]">300</div>
  </div>
</template>

<style lang="scss" scoped>
:deep(.ProseMirror) {
  &.ProseMirror-focused {
    outline: none !important;
  }

  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    font-weight: bold;
  }

  h1 {
    font-size: 36px;
  }
  h2 {
    font-size: 32px;
  }
  h3 {
    font-size: 28px;
  }
  h4 {
    font-size: 24px;
  }
  h5 {
    font-size: 20px;
  }
  h6 {
    font-size: 16px;
  }

  ul {
    list-style-type: disc;
    list-style-position: inside;

    li {
      * {
        display: inline-block;
      }
    }
  }

  ol {
    list-style-type: decimal;
    list-style-position: inside;

    li {
      * {
        display: inline-block;
      }
    }
  }
}
</style>
