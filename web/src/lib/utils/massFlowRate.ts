export function computeMassFlowRate(oxFlow: number[], timestep: number): number[] {
    if (oxFlow.length < 2) return [];

    const rates: number[] = [];
    for (let i = 1; i < oxFlow.length; i++) {
        const rate = (oxFlow[i] - oxFlow[i - 1]) / timestep;
        rates.push(rate);
    }

    return rates;
}
