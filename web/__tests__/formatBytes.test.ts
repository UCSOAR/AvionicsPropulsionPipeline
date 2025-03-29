import { formatBytes } from "../src/lib/utils/fetchUsage"

describe("formatBytes", () => {
  test("returns '0 Bytes' when input is 0", () => {
    expect(formatBytes(0)).toBe("0 Bytes");
  });

  test("formats bytes correctly", () => {
    expect(formatBytes(1024)).toBe("1.00 KB");
    expect(formatBytes(1048576)).toBe("1.00 MB");
    expect(formatBytes(1073741824)).toBe("1.00 GB");
  });

  test("handles decimal precision correctly", () => {
    expect(formatBytes(1536, 1)).toBe("1.5 KB");
    expect(formatBytes(1536, 3)).toBe("1.500 KB");
  });

  test("returns bytes as-is if value is below 1 KB", () => {
    expect(formatBytes(512)).toBe("512 Bytes");
  });
});
