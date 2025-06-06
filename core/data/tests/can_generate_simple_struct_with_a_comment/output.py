from __future__ import annotations

from pydantic import BaseModel, Field
from typing import List, Optional


class Location(BaseModel):
    pass
class Person(BaseModel):
    """
    This is a comment.
    """
    name: str
    """
    This is another comment
    """
    age: int
    info: Optional[str] = Field(default=None)
    emails: List[str]
    location: Location

