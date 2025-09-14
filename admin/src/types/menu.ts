interface MenuData {
  id: number;
  parent_id: number;
  name: string;
  path: string;
  icon: string;
  create_at: string;
  update_at: string;
  state: number;
}

interface MenuTreeData extends MenuData {
  children?: MenuTreeData[];
}

export type { MenuData, MenuTreeData };
