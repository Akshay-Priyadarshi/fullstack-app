import { useState } from "react";

export const useCounter = (initialCount: number = 0) => {
  const [count, setCount] = useState(initialCount);

  const increase = () => {
    setCount((prevCount) => prevCount + 1);
  };

  const decrease = () => {
    setCount((prevCount) => prevCount - 1);
  };

  const reset = () => {
    setCount(initialCount);
  };

  return { count, increase, decrease, reset };
};
