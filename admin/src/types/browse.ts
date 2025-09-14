interface BrowseCountData {
  name: string;
  date: string;
  trend: string;
  count: number;
  label: string[];
  value: number[];
  group: {
    label: string;
    value: number;
  }[];
}

interface BrowseRegionData {
  rank: number;
  name: string;
  value: number;
}

export type { BrowseCountData, BrowseRegionData };
