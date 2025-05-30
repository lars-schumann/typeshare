from __future__ import annotations

from pydantic import BaseModel, Field
from typing import Optional


class Foo(BaseModel):
    bar: Optional[bool] = Field(default=None)

