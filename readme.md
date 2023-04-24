# MiMeMail

Initial docs

## Optimizations

That was interesting, never thought that checking the validation of JSON would be so "expensive", one of the ways to optimize the code but, wouldnt be safe at all, is if i dont unmarshal and copy it directly to an interface.

Other data that i found

### Searchs with max results 20, highlight on, 14.5k "hey" to the API

<table>
<thead>
<tr>
<th>Word</th>
<th>Value (Return)</th>
<th>Runtime</th>
</tr>
</thead>
<tbody align="center">
<tr><td>'example'</td><td>11288</td><td>20 ms</td></tr>
<tr><td>'mutation'</td><td>9</td><td>3 ms</td></tr>
<tr><td>'mutation, where are you'</td><td>3</td><td>3 ms</td></tr>
<tr><td>'how are you'</td><td>1818</td><td>250 ms</td></tr>
<tr><td>'how are'</td><td>3388</td><td>150 ms</td></tr>
</tbody>
</table>
