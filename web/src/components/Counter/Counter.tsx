import { useCounter } from "@hooks";
import { Button } from "../ui";

export interface ICounterProps {
  initialCount?: number;
}

export const Counter: React.FC<ICounterProps> = ({ initialCount }) => {
  const { count, increase, decrease, reset } = useCounter(initialCount ?? 0);

  return (
    <div className="flex flex-col gap-4 items-center bg-slate-400 p-4 rounded-lg shadow-lg">
      <h1
        data-testid="counter-count"
        className="text-4xl"
      >
        {count}
      </h1>
      <div className="flex gap-4">
        <Button onClick={increase}>Increase</Button>
        <Button onClick={reset}>Reset</Button>
        <Button onClick={decrease}>Decrease</Button>
      </div>
    </div>
  );
};
