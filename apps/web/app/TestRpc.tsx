"use client";

import { client } from "../lib/rpc/client";
import { use } from "react";

const response = client.sayHello({ name: "Skibidi", age: 1 });

export function TestRpc() {
  const result = use(response);
  return <div>{JSON.stringify(result, null, 2)}</div>;
}
