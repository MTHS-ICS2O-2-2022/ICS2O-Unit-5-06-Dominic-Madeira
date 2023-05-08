// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"


/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {
  // input
  const numberA = parseInt(document.getElementById('number-a').value)
  const numberB = parseInt(document.getElementById('number-b').value)
  let counter = 0
  let output = 0

  // process


  while (counter < numberA) {
    output = output + numberB
    counter = counter + 1
  }
  document.getElementById('answer').innerHTML = numberA + " x " + numberB + " = " + output
}