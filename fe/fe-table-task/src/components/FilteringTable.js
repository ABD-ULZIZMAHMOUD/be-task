import React, { useMemo, useState, useEffect } from "react";
import { useTable, usePagination } from "react-table";
import { COLUMNS } from "./columns";
import "./table.css";
import { SelectInput } from "./SelectInput";
import axios from "axios";

export const FilteringTable = () => {

  const [data, setData] = useState([]);

  const [isLoading, setIsLoading] = useState(true);
  const [country, setCountry] = useState("");
  const [stateFilter, setStateFilter] = useState("");
  useEffect(() => {
    axios.get("http://127.0.0.1:8000/customers/all?limit=50&page=1&" + "state=" + stateFilter + "&country=" + country).then((res) => {
      setData(res.data.payload.records);
      setIsLoading(false);
    });

  }, [country, stateFilter]);

  function changeHandlerState(event) {
    setStateFilter(event==null?"":event.value)
  }

  function handleSelectChange(event) {
    setCountry(event==null?"":event.value);
  }
  console.log("data", data);
  const columns = useMemo(() => COLUMNS, []);
  const tableInstance = useTable(
    {
      columns, // columns: columns,
      data,
    },
    usePagination
  );
  const {
    getTableBodyProps,
    getTableProps,
    headerGroups,
    page,
    nextPage,
    previousPage,
    canNextPage,
    pageOptions,
    canPreviousPage,
    prepareRow,
    state,
  } = tableInstance;
  const { pageIndex } = state;
  const stateOptions = [
    { value: "valid", label: "Valid" },
    { value: "not_valid", label: "Not Valid" },
  ];

  const countryOptions = [
    { value: "Cameroon", label: "Cameroon" },
    { value: "Ethiopia", label: "Ethiopia" },
    { value: "Morocco", label: "Morocco" },
    { value: "Mozambique", label: "Mozambique" },
    { value: "Uganda", label: "Uganda" },

  ];
  if (isLoading) {
    return <h2> loading</h2>
  }
  return (
    <>
      <div
        style={{
          display: "flex",

          justifyContent: "space-between",
          margin: "30px 50px",
        }}
      >
        <div
          style={{
            display: "flex",
            gap: 40,
          }}
        >
          <SelectInput options={countryOptions} placeholder="select countries" setSelectedOption={handleSelectChange } selectedOption={country} />
          <SelectInput options={stateOptions} placeholder="select state" setSelectedOption={changeHandlerState} selectedOption={stateFilter} />
        </div>
        <div
          style={{
            display: "flex",
            gap: 40,
          }}
        >
        </div>
      </div>
      <table {...getTableProps()}>
        <thead>
          {headerGroups.map((headerGroup) => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map((col) => (
                <th {...col.getHeaderProps()}>{col.render("Header")}</th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {page.map((row) => {
            prepareRow(row);
            return (
              <tr {...row.getRowProps()}>
                {row.cells.map((cell) => (
                  <td {...cell.getCellProps()}>{cell.render("Cell")}</td>
                ))}
              </tr>
            );
          })}
        </tbody>
      </table>
      <div>
        <button onClick={() => previousPage()} disabled={!canPreviousPage}>
          {" "}
          {`<`}{" "}
        </button>
        <strong>{` page ${pageIndex + 1} of ${pageOptions.length} `}</strong>
        <button onClick={() => nextPage()} disabled={!canNextPage}>
          {" "}
          {`>`}{" "}
        </button>
      </div>
    </>
  );
};
