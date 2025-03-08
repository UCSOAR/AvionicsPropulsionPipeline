"use client";

import React, { useEffect, useState } from "react";
import Plot from "react-plotly.js";
import { useTheme } from "next-themes";
import { useMetadataStore } from "@/stores/metadataStore";

/**
 * A Plotly-based line chart that resembles Recharts styling,
 * using WebGL (scattergl) for performance.
 */
export function LineChart() {
  const { resolvedTheme } = useTheme();
  const isDark = resolvedTheme === "dark";

  const metadataStore = useMetadataStore();
  const { metadata, colX, colY } = metadataStore;

  // Local state for the processed chart data
  const [chartData, setChartData] = useState<{ x: number; y: number }[]>([]);

  useEffect(() => {
    if (!metadata || !colX || !colY) {
      setChartData([]);
      return;
    }

    // Check if the actual data for colX & colY is present
    const xInfo = metadata.xColumns?.[colX];
    const yInfo = metadata.yColumns?.[colY];

    if (!xInfo || !yInfo) {
      // Data hasn't been fetched yet, or columns are missing
      setChartData([]);
      return;
    }

    const xValues = xInfo.rows ?? [];
    const yValues = yInfo.rows ?? [];

    if (xValues.length !== yValues.length) {
      console.warn("Mismatch in X and Y data lengths.");
      setChartData([]);
      return;
    }

    // Combine and sort by X
    const combined = xValues.map((x, i) => ({ x, y: yValues[i] }));
    combined.sort((a, b) => a.x - b.x);

    setChartData(combined);
  }, [metadata, colX, colY]);

  // If no data, show a placeholder
  if (!chartData.length) {
    return (
      <div className="flex items-center justify-center h-full text-sm text-gray-500">
        No Plotly data yet...
      </div>
    );
  }

  // Prepare the Plotly data & layout
  const plotlyData: Partial<Plotly.Data>[] = [
    {
      x: chartData.map((d) => d.x),
      y: chartData.map((d) => d.y),
      type: "scattergl", // WebGL-based scatter
      mode: "lines",
      marker: { size: 6 },
      line: { width: 2, color: isDark ? "#ff5a5a" : "#dc2626" }, // A red shade
    },
  ];

  // Replicate Recharts-like styling
  const layout: Partial<Plotly.Layout> = {
    dragmode: "pan",
    plot_bgcolor: isDark ? "#1e1e1e" : "#ffffff", // Chart background
    paper_bgcolor: isDark ? "#1e1e1e" : "#ffffff", // Outer background
    font: {
      color: isDark ? "#f8fafc" : "#0f172a", // Axis/legend text
    },
    xaxis: {
      title: "X Axis",
      gridcolor: isDark ? "#334155" : "#e2e8f0",
    },
    yaxis: {
      title: "Y Axis",
      gridcolor: isDark ? "#334155" : "#e2e8f0",
    },
    margin: { l: 50, r: 30, t: 30, b: 50 },
    showlegend: false, // or true if you want a legend
  };

  const config: Partial<Plotly.Config> = {
    responsive: true,
    displayModeBar: false, // Hide the Plotly toolbar
  };

  return (
    <Plot
      data={plotlyData}
      layout={layout}
      config={config}
      style={{ width: "100%", height: "100%" }}
      useResizeHandler
    />
  );
}
