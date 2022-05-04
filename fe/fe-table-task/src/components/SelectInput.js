import { useState } from "react";
import Select from "react-select";

export const SelectInput = ({ options, placeholder,setSelectedOption,selectedOption }) => {
 // const [selectedOption] = useState(null);
  return (
    <div style={{ width: 200 }}>
      <Select
        defaultValue={selectedOption}
        onChange={setSelectedOption}
        options={options}
        isClearable={true}
        placeholder={placeholder}
      />
    </div>
  );
};
