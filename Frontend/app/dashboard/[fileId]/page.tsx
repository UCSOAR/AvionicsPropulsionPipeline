"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useMetadataStore } from "@/stores/metadataStore";
import { endpointMapping } from "@/utils/constants"; // Same as in Vue
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { LineChart } from "@/components/charts/line-chart";
import { BarChart } from "@/components/charts/bar-chart";
import { PieChart } from "@/components/charts/pie-chart";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface DashboardProps {
  params: {
    fileId: string;
  };
}

export default function Dashboard({ params }: DashboardProps) {
  const router = useRouter();
  const metadataStore = useMetadataStore();

  // Loading simulation
  const [loading, setLoading] = useState(true);
  const [progress, setProgress] = useState(0);

  // Column names from metadata
  const xColumnNames = metadataStore.metadata?.xColumnNames ?? [];
  const yColumnNames = metadataStore.metadata?.yColumnNames ?? [];

  // The currently selected X & Y columns
  const [selectedXValue, setSelectedXValue] = useState("");
  const [selectedYValue, setSelectedYValue] = useState("");

  // ▶ State for computed metrics (Min, Max, Avg)
  const [minValue, setMinValue] = useState<number | null>(null);
  const [maxValue, setMaxValue] = useState<number | null>(null);
  const [avgValue, setAvgValue] = useState<number | null>(null);

  // ─────────────────────────────────────────────────────────────
  // 1) Simulate a loading spinner (like your Vue code might have)
  // ─────────────────────────────────────────────────────────────
  useEffect(() => {
    const timer1 = setTimeout(() => setProgress(30), 500);
    const timer2 = setTimeout(() => setProgress(60), 1000);
    const timer3 = setTimeout(() => {
      setProgress(100);
      setLoading(false);
    }, 1500);
    return () => {
      clearTimeout(timer1);
      clearTimeout(timer2);
      clearTimeout(timer3);
    };
  }, []);

  // ─────────────────────────────────────────────────────────────
  // 2) Auto-select first columns & fetch data for them
  // ─────────────────────────────────────────────────────────────
  useEffect(() => {
    if (xColumnNames.length > 0 && yColumnNames.length > 0) {
      const defaultX = xColumnNames[0];
      const defaultY = yColumnNames[0];
      if (!selectedXValue) setSelectedXValue(defaultX);
      if (!selectedYValue) setSelectedYValue(defaultY);
      fetchColumnData(defaultX, defaultY);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [xColumnNames, yColumnNames]);

  // ─────────────────────────────────────────────────────────────
  // 3) Post the same payload to fetch actual column rows
  // ─────────────────────────────────────────────────────────────
  async function fetchColumnData(colX: string, colY: string) {
    try {
      const name = metadataStore.name;
      if (!name || !colX || !colY) {
        console.warn("Cannot fetch columns: missing name or colX/colY");
        return;
      }

      // Same payload as in Vue
      const payload = {
        name,
        startRow: 0,        // or your actual start row
        numRows: 10000,     // or however many you want
        xColumnNames: [colX],
        yColumnNames: [colY],
      };

      console.log("Posting to getStaticFireColumnsUrl:", payload);

      const response = await fetch(endpointMapping.getStaticFireColumnsUrl.toString(), {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        throw new Error("Failed to fetch column data");
      }

      // The response shape should match your Vue code
      const data = await response.json();
      console.log("Received column data:", data);

      // Merge the returned columns into metadata
      metadataStore.setMetadata(
        {
          ...metadataStore.metadata,
          xColumns: {
            ...(metadataStore.metadata?.xColumns || {}),
            ...(data.xColumns || {}),
          },
          yColumns: {
            ...(metadataStore.metadata?.yColumns || {}),
            ...(data.yColumns || {}),
          },
        },
        colX,
        colY,
        name
      );
    } catch (err) {
      console.error("Error fetching column data:", err);
    }
  }

  // ─────────────────────────────────────────────────────────────
  // 4) If user selects different X/Y, fetch new columns
  // ─────────────────────────────────────────────────────────────
  function handleSelectionChange(newX: string, newY: string) {
    setSelectedXValue(newX);
    setSelectedYValue(newY);
    fetchColumnData(newX, newY);
  }

  // ─────────────────────────────────────────────────────────────
  // 5) Compute min, max, avg from the selected Y column
  // ─────────────────────────────────────────────────────────────
  useEffect(() => {
    const yRows = metadataStore.metadata?.yColumns?.[selectedYValue]?.rows ?? [];
    if (yRows.length === 0) {
      setMinValue(null);
      setMaxValue(null);
      setAvgValue(null);
      return;
    }

    const min = Math.min(...yRows);
    const max = Math.max(...yRows);
    const avg = yRows.reduce((acc: number, val: number) => acc + val, 0) / yRows.length;

    setMinValue(min);
    setMaxValue(max);
    setAvgValue(avg);
  }, [metadataStore.metadata, selectedYValue]);

  // ─────────────────────────────────────────────────────────────
  // 6) Render the dashboard
  // ─────────────────────────────────────────────────────────────
  return (
    <div className="flex flex-col gap-6">
      {/* Header Section */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 className="text-2xl font-bold">
            Dashboard for {decodeURIComponent(params.fileId)}
          </h1>
          <p className="text-muted-foreground">
            Visualizing data for {selectedYValue}
          </p>
        </div>

        {/* X/Y Selection - same styling */}
        <div className="flex flex-wrap gap-2 p-4 rounded-lg">
          <div className="flex flex-col">
            <label className="text-sm font-medium text-gray-300">X Column:</label>
            <Select
              value={selectedXValue}
              onValueChange={(val) => handleSelectionChange(val, selectedYValue)}
            >
              <SelectTrigger className="w-[150px] bg-black text-white border-none focus:ring-0">
                <SelectValue placeholder="Select X Column" />
              </SelectTrigger>
              <SelectContent className="bg-black text-white border-none">
                {xColumnNames.map((col) => (
                  <SelectItem key={col} value={col} className="text-white hover:bg-gray-700">
                    {col}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="flex flex-col">
            <label className="text-sm font-medium text-gray-300">Y Column:</label>
            <Select
              value={selectedYValue}
              onValueChange={(val) => handleSelectionChange(selectedXValue, val)}
            >
              <SelectTrigger className="w-[150px] bg-black text-white border-none focus:ring-0">
                <SelectValue placeholder="Select Y Column" />
              </SelectTrigger>
              <SelectContent className="bg-black text-white border-none">
                {yColumnNames.map((col) => (
                  <SelectItem key={col} value={col} className="text-white hover:bg-gray-700">
                    {col}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>
      </div>

      {/* Loading Spinner */}
      {loading ? (
        <div className="flex flex-col items-center justify-center h-[80vh] gap-4">
          <h2 className="text-2xl font-bold">Loading file data...</h2>
          <Progress value={progress} className="w-[60%] max-w-md" />
        </div>
      ) : (
        <div className="space-y-6">
          {/* Metric Cards (Min, Max, Avg) */}
          <div className="grid gap-4 md:grid-cols-3">
            <Card className="shadow-sm hover:shadow transition-shadow">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium">Minimum Value</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">
                  {minValue !== null ? minValue.toFixed(2) : "N/A"}
                </div>
              </CardContent>
            </Card>

            <Card className="shadow-sm hover:shadow transition-shadow">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium">Maximum Value</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">
                  {maxValue !== null ? maxValue.toFixed(2) : "N/A"}
                </div>
              </CardContent>
            </Card>

            <Card className="shadow-sm hover:shadow transition-shadow">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium">Average Value</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">
                  {avgValue !== null ? avgValue.toFixed(2) : "N/A"}
                </div>
              </CardContent>
            </Card>
          </div>

          {/* Line Chart */}
          <div className="grid gap-4 md:grid-cols-2">
            <Card className="col-span-2">
              <CardHeader>
                <CardTitle>Static Fire Chart for {selectedYValue}</CardTitle>
              </CardHeader>
              <CardContent className="h-[300px]">
                <LineChart />
              </CardContent>
            </Card>
          </div>

          {/* Bar & Pie Charts (Same as your code) */}
          <div className="grid gap-4 md:grid-cols-2">
            <Card>
              <CardHeader>
                <CardTitle>Revenue by Category</CardTitle>
              </CardHeader>
              <CardContent className="h-[300px]">
                <BarChart />
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>Traffic Sources</CardTitle>
              </CardHeader>
              <CardContent className="h-[300px]">
                <PieChart />
              </CardContent>
            </Card>
          </div>
        </div>
      )}
    </div>
  );
}
