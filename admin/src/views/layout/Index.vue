<script setup lang="ts">
import { ref, watch, onMounted } from "vue";
import { useRoute } from "vue-router";
import { useDark, useToggle } from "@vueuse/core";
import { Moon, Sunny } from "@element-plus/icons-vue";
import { reqMenuList } from "@/api";
import type { ResponseData, MenuData, MenuTreeData } from "@/types";

const route = useRoute();
const isDark = useDark();
const toggleDark = useToggle(isDark);
const menuData = ref<MenuTreeData[]>([]);
const breadcrumbData = ref<MenuData[]>([]);

// 简洁优雅的面包屑查找方法
const findBreadcrumb = (items: MenuTreeData[], currentPath: string) => {
  // 重置面包屑数据
  breadcrumbData.value.length = 0;

  // 内部递归函数，专注于路径查找逻辑
  const findPath = (
    menus: MenuTreeData[],
    targetPath: string,
    parentChain: MenuData[] = []
  ): boolean => {
    for (const menu of menus) {
      const fullPath = `${menu.path}/${route.path.split(`/`).pop()}`;

      // 找到匹配路径时，构建完整面包屑链
      if (fullPath === targetPath) {
        // 添加父菜单链和当前菜单
        breadcrumbData.value = [...parentChain, menu];
        return true;
      }

      // 有子菜单时递归查找，并传递当前菜单作为父链
      if (menu.children?.length) {
        if (findPath(menu.children, targetPath, [...parentChain, menu])) {
          return true;
        }
      }
    }
    return false;
  };

  // 开始查找
  findPath(items, currentPath);
};

const arrayToTree = (data: MenuData[]): MenuTreeData[] => {
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

const getMenuList = async () => {
  const response: ResponseData<MenuData[]> = await reqMenuList();
  if (response.code === 200) {
    menuData.value = arrayToTree(response.data!);
    findBreadcrumb(menuData.value, route.path);
  }
};

const getRealPath = () => {
  return route.path.split('/').slice(0, -1).join('/');
};

onMounted(async () => {
  await getMenuList();
});

watch(
  () => route.path,
  (path) => {
    findBreadcrumb(menuData.value, path);
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
          <template v-for="value in menuData" :key="value.id">
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
                v-for="value in breadcrumbData"
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
