<template>
  <div ref="chartRef" :class="className"></div>
</template>

<script lang="ts" setup>
import * as echarts from "echarts";
import { onMounted, onUnmounted, ref } from "vue";
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

const chartRef = ref<HTMLElement | null>(null);

let chartInstance: echarts.ECharts | null = null;

const onResize = debounce(() => {
  chartInstance?.resize();
}, 300);

const initChart = () => {
  chartInstance = echarts.init(chartRef.value);

  const chartOpt: echarts.EChartsOption = {
    tooltip: {
      trigger: "axis",
    },
    xAxis: {
      type: "category",
      data: props.xAxisData as string[],
      show: false,
      boundaryGap: false,
    },
    grid: {
      left: 0,
      right: 0,
      bottom: 0,
      top: 0,
    },
    yAxis: {
      type: "value",
      show: false,
    },
    series: [
      {
        name: "访问量",
        type: "line",
        symbol: "none",
        sampling: "lttb",
        itemStyle: {
          color: "#409eff",
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: "rgba(64, 158, 255, 0.8)",
            },
            {
              offset: 1,
              color: "rgba(64, 158, 255, 0.2)",
            },
          ]),
        },
        lineStyle: {
          width: 2,
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
  chartInstance?.dispose();
  chartInstance = null;
  window.removeEventListener("resize", onResize);
});
</script>

<style lang="scss" scoped></style>
