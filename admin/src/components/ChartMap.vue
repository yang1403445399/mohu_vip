<template>
  <div ref="chartRef" :class="className"></div>
</template>

<script lang="ts" setup>
import * as echarts from "echarts";
import { useDark } from "@vueuse/core";
import { onMounted, onUnmounted, ref, watch } from "vue";
import { debounce } from "lodash";
import type { BrowseRegionData } from "@/types/browse";

const props = defineProps({
  className: {
    type: String,
    default: "",
  },
  maxValue: {
    type: Number,
    default: 0,
  },
  geoData: {
    type: Object,
    default: () => {},
  },
  mapData: {
    type: Array,
    default: () => [],
  },
  barData: {
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
  echarts.registerMap("china", props.geoData as any);

  chartInstance = echarts.init(chartRef.value);

  const chartOpt: echarts.EChartsOption = {
    dataset: {
      source: props.barData as (string | number)[][],
    },
    tooltip: {
      trigger: "item",
    },
    visualMap: {
      left: "right",
      min: 0,
      max: props.maxValue > 0 ? props.maxValue : 1000000,
      inRange: {
        color: isDark.value ? ["#303643", "#409eff"] : ["#f9f9f9", "#409eff"],
      },
      text: ["高", "低"],
      textStyle: {
        color: isDark.value ? "#ffffff" : "#606266",
      },
      calculable: true,
    },
    grid: {
      top: 10,
      bottom: 0,
      left: 0,
      right: 0,
      width: "18%",
    },
    xAxis: {
      name: "value",
      show: false,
    },
    yAxis: {
      type: "category",
      show: false,
    },
    series: [
      {
        name: "中华人民共和国",
        map: "china",
        type: "map",
        roam: false,
        zoom: 1,
        center: ["50%", "50%"],
        itemStyle: {
          borderColor: "#409eff",
          borderWidth: 2,
        },
        emphasis: {
          label: {
            show: false,
          },
          itemStyle: {
            borderColor: "#409eff",
            borderWidth: 2,
            areaColor: "#409eff",
          },
        },
        data: props.mapData as BrowseRegionData[],
      },
      {
        name: "中华人民共和国",
        type: "bar",
        encode: {
          x: "value",
          y: "name",
        },
        barWidth: "28%",
        showBackground: true,
        label: {
          show: true,
          position: [0, -16],
          formatter: (params: any) => {
            return `NO:${params.value[0]} | ${params.value[1]} | ${params.value[2]}`;
          },
        },
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
      visualMap: {
        inRange: {
          color: newVal ? ["#303643", "#409eff"] : ["#f9f9f9", "#409eff"],
        },
        textStyle: {
          color: newVal ? "#ffffff" : "#606266",
        },
      },
    })
  }
});
</script>

<style lang="scss" scoped>
/* 地图组件样式 */
</style>
