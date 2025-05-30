from __future__ import annotations

from pydantic import BaseModel, ConfigDict, Field
from typing import Dict, List, Optional


class CustomType(BaseModel):
    pass
class Types(BaseModel):
    model_config = ConfigDict(populate_by_name=True)

    s: str
    static_s: str
    int_8: int = Field(alias="int8")
    float: float
    double: float
    array: List[str]
    fixed_length_array: List[str]
    dictionary: Dict[str, int]
    optional_dictionary: Optional[Dict[str, int]] = Field(default=None)
    custom_type: CustomType

