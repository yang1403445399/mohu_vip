<template>
  <div ref="chartRef" :class="className"></div>
</template>

<script lang="ts" setup>
import * as echarts from "echarts";
import { useDark } from "@vueuse/core";
import { onMounted, onUnmounted, ref, watch } from "vue";
import { debounce } from "lodash";

const props = defineProps({
  className: {
    type: String,
    default: "",
  },
  yAxisData: {
    type: Array,
    default: () => [],
  },
  xAxisData: {
    type: Array,
    default: () => [],
  },
});

const isDark = useDark();

const chartRef = ref<HTMLElement | null>(null);

let chartInstance: echarts.ECharts | null = null;

const onResize = debounce(() => {
  chartInstance?.resize();
}, 300);

const onDispose = () => {
  chartInstance?.dispose();
  chartInstance = null;
};

const initChart = () => {
  chartInstance = echarts.init(chartRef.value);

  const chartOpt: echarts.EChartsOption = {
    tooltip: {
      trigger: "axis",
    },
    xAxis: {
      type: "category",
      data: props.xAxisData as string[],
      axisLine: {
        lineStyle: {
          color: isDark.value ? "#4c4d4f" : "#606266",
        },
      },
      axisLabel: {
        color: isDark.value ? "#ffffff" : "#606266",
      },
    },
    grid: {
      left: 0,
      right: 0,
      bottom: 0,
      top: 0,
    },
    yAxis: {
      type: "value",
      axisLabel: {
        color: isDark.value ? "#ffffff" : "#606266",
      },
      splitLine: {
        lineStyle: {
          color: isDark.value ? "#4c4d4f" : "#e3e3e3",
        },
      },
    },
    series: [
      {
        name: "发布量",
        type: "bar",
        showBackground: true,
        itemStyle: {
          color: "#409eff",
        },
        data: props.yAxisData as number[],
      },
    ],
  };

  chartInstance.setOption(chartOpt);
};

onMounted(() => {
  initChart();
  window.addEventListener("resize", onResize);
});

onUnmounted(() => {
  onDispose();
  window.removeEventListener("resize", onResize);
});

watch(isDark, (newVal) => {
  if (chartInstance) {
    chartInstance.setOption({
      xAxis: {
        axisLine: {
          lineStyle: {
            color: newVal ? "#4c4d4f" : "#606266",
          },
        },
        axisLabel: {
          color: newVal ? "#ffffff" : "#606266",
        },
      },
      yAxis: {
        axisLabel: {
          color: newVal ? "#ffffff" : "#606266",
        },
        splitLine: {
          lineStyle: {
            color: newVal ? "#4c4d4f" : "#e3e3e3",
          },
        },
      },
    });
  }
});
</script>

<style lang="scss" scoped>
/* 地图组件样式 */
</style>
