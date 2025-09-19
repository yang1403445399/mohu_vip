const trimParser = (value: string | undefined | null): string => {
  return value?.trim() || "";
};

const removeAllSpaces = (value: string | undefined | null): string => {
  return value?.replace(/\s+/g, "") || "";
};

const toLowerCase = (value: string | undefined | null): string => {
  return value?.toLowerCase() || "";
};

const toUpperCase = (value: string | undefined | null): string => {
  return value?.toUpperCase() || "";
};

export {
  trimParser,
  removeAllSpaces,
  toLowerCase,
  toUpperCase,
};
