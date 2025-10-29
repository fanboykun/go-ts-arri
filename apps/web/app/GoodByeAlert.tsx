"use client";

import { useCallback } from "react";
import { client } from "../lib/rpc/client";
import styles from "./page.module.css";

export function GoodByeAlert() {
  const sayGoodBye = useCallback(async () => {
    const res = await client.sayGoodbye({ name: "namamu", age: 0 });
    alert(JSON.stringify(res, null, 2));
  }, []);

  const getAddress = useCallback(async () => {
    const res = await client.getAddress({ address: null });
    alert(JSON.stringify(res, null, 2));
  }, []);

  const getUsers = useCallback(async () => {
    const res = await client.getUsers({
      limit: 10,
      offset: 1,
      page: 1,
    });
    console.log(res);
    alert("check log for response");
  }, []);
  return (
    <div>
      <button type="button" className={styles.secondary} onClick={getUsers}>
        Dummy Pagination
      </button>
      <button type="button" className={styles.secondary} onClick={getAddress}>
        Address
      </button>
      <button type="button" className={styles.secondary} onClick={sayGoodBye}>
        Goodbye
      </button>
    </div>
  );
}
