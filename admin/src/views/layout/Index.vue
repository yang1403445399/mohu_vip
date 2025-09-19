<script setup lang="ts">
import { reactive, watch, onMounted } from "vue";
import { useRoute } from "vue-router";
import { useDark, useToggle } from "@vueuse/core";
import { Moon, Sunny } from "@element-plus/icons-vue";
import { reqMenuList } from "@/api/menu";
import type { ResponseData } from "@/types/common";
import type { MenuData, MenuTreeData } from "@/types/menu";

const route = useRoute();
const isDark = useDark();
const toggleDark = useToggle(isDark);
const menuData = reactive({
  list: [] as MenuData[],
  tree: [] as MenuTreeData[],
  breadcrumb: [] as MenuData[],
});

const getMenuList = async () => {
  const response: ResponseData<MenuData[]> = await reqMenuList();
  if (response.code === 200) {
    menuData.list = response.data!;
    menuData.tree = buildNavigationMenuTree(response.data!);
    findBreadcrumbItems(menuData.list, route.path);
  }
};

const buildNavigationMenuTree = (data: MenuData[]): MenuTreeData[] => {
  const menuTree: MenuTreeData[] = [];
  const menuMap = new Map<number, MenuTreeData>();

  data.forEach((item) => {
    const menuItem = { ...item, children: [] } as MenuTreeData;
    menuMap.set(item.id, menuItem);
  });

  data.forEach((item) => {
    const menuItem = menuMap.get(item.id)!;

    if (item.parent_id === 0) {
      menuTree.push(menuItem);
    } else {
      const parent = menuMap.get(item.parent_id);
      if (parent) {
        parent.children!.push(menuItem);
      } else {
        menuTree.push(menuItem);
      }
    }
  });

  return menuTree;
};

const findBreadcrumbItems = (items: MenuData[], currentPath: string) => {
  menuData.breadcrumb = [];

  const currentItem = items.find(
    (item) => `${item.path}/${route.path.split(`/`).pop()}` === currentPath
  );

  if (currentItem) {
    menuData.breadcrumb.push(currentItem);

    let parentId = currentItem.parent_id;
    while (parentId !== 0) {
      const parentItem = items.find((item) => item.id === parentId);
      if (parentItem) {
        menuData.breadcrumb.unshift(parentItem);
        parentId = parentItem.parent_id;
      } else {
        break;
      }
    }
  }
};

const getRealPath = () => {
  return route.path.split("/").slice(0, -1).join("/");
};

const onInit = async () => {
  await getMenuList();
};

onMounted(async () => {
  await onInit();
});

watch(
  () => route.path,
  (path) => {
    findBreadcrumbItems(menuData.list, path);
  }
);
</script>

<template>
  <el-container class="h-screen">
    <el-header class="!h-auto border-b border-[var(--el-border-color)]">
      <div class="w-330 m-auto flex justify-between items-center">
        <div class="w-50 h-15 flex items-center">
          <img class="w-30" src="@/assets/logo_002.svg" alt="" />
        </div>
        <div class="flex items-center">
          <div class="flex mr-4">
            <el-switch
              v-model="isDark"
              :active-action-icon="Moon"
              :inactive-action-icon="Sunny"
              @change="toggleDark"
            />
          </div>
          <el-dropdown>
            <div class="flex items-center outline-none cursor-pointer">
              <el-avatar
                :size="28"
                src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
              />
              <div
                class="ml-2 mr-1 text-[var(--el-text-color-primary)] leading-none text-sm"
              >
                ADMIN
              </div>
              <el-icon color="var(--el-text-color-primary)" :size="16"
                ><CaretBottom
              /></el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>Action 1</el-dropdown-item>
                <el-dropdown-item>Action 2</el-dropdown-item>
                <el-dropdown-item>Action 3</el-dropdown-item>
                <el-dropdown-item disabled>Action 4</el-dropdown-item>
                <el-dropdown-item divided>Action 5</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
      <div class="border-t border-[var(--el-border-color)] w-330 m-auto">
        <el-menu
          :default-active="getRealPath()"
          mode="horizontal"
          class="!border-none"
          router
        >
          <template v-for="value in menuData.tree" :key="value.id">
            <el-menu-item
              :index="value.path"
              v-if="value.children!.length === 0"
            >
              <el-icon><component :is="value.icon" /></el-icon>
              <template #title>{{ value.name }}</template>
            </el-menu-item>
            <el-sub-menu index="/admin/setup" v-if="value.children!.length > 0">
              <template #title>
                <el-icon><component :is="value.icon" /></el-icon>
                <span>{{ value.name }}</span>
              </template>
              <template v-for="child in value.children" :key="child.id">
                <el-menu-item :index="child.path">{{
                  child.name
                }}</el-menu-item>
              </template>
            </el-sub-menu>
          </template>
        </el-menu>
      </div>
    </el-header>
    <el-main class="!p-0 bg-[var(--grey-bg-color)]">
      <el-scrollbar wrap-class="h-full">
        <div class="w-330 m-auto py-4">
          <div class="pb-4">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item :to="{ path: '/admin' }"
                ><el-icon><House /></el-icon
              ></el-breadcrumb-item>
              <el-breadcrumb-item
                v-for="value in menuData.breadcrumb"
                :key="value.id"
                :to="{ path: value.path }"
                >{{ value.name }}</el-breadcrumb-item
              >
            </el-breadcrumb>
          </div>
          <router-view />
        </div>
      </el-scrollbar>
    </el-main>
    <el-footer class="!h-10 border-t border-[var(--el-border-color)]">
      <div
        class="text-[#999999] leading-none text-xs flex items-center justify-center h-full"
      >
        2025 © Mohu 版权所有
      </div>
    </el-footer>
  </el-container>
</template>

<style lang="scss" scoped>
:deep(.el-menu) {
  height: 54px;

  .el-menu-item {
    height: 100%;

    &:not(.is-disabled) {
      &:hover {
        background: none;
      }

      &:focus {
        background: none;
        color: var(--el-menu-text-color);
      }
    }

    &.is-active {
      background: none;
    }
  }
}
</style>
