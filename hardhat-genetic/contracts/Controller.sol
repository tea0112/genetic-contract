// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract Controller {
  struct Movie {
        uint256 id;
        string title;
        uint256 length;
        // address createdBy;
  }

  mapping(uint256 => Movie) public movies;

  event CreatingMovieEvent(string message, uint256 movieId);

  function createNewMovie(uint256 _id, string memory _title, uint256 _length) external {
    movies[_id] = Movie({id: _id, title: _title, length: _length});
    emit CreatingMovieEvent("This is the message of creating movie event", _id);
  }

  function getMovie(uint256 _id) external view returns (Movie memory) {
    return movies[_id];
  }

  function deleteMovie(uint256 _id) public {
    delete movies[_id];
  }
}