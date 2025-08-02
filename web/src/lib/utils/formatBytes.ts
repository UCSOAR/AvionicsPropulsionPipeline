// Function to Format Bytes
export function formatBytes(bytes: number, decimals: number = 2): string {
  if (bytes === 0) return "0 Bytes";

  const sizes = ["Bytes", "KB", "MB", "GB"];
  const i = Math.floor(Math.log(bytes) / Math.log(1024));

  if (i >= sizes.length) return `${bytes} Bytes`;

  const value = bytes / Math.pow(1024, i);
  return i === 0 || value < 1
    ? `${Math.floor(value)} ${sizes[i]}`
    : `${value.toFixed(decimals)} ${sizes[i]}`;
}
