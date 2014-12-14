package btcnet
import "github.com/hlandauf/btcwire"

type BugType int
const (
  BugNone BugType = iota
  BugFullyIgnore
  BugFullyApply
  BugInUTXO
  BugBitcoin
)

// Determine if a transaction represents a historic bug. This currently considers all
// supported currencies so as to make it context-free.
func IsHistoricBug(txHeight int64, txId *btcwire.ShaHash) (bugType BugType, isBug bool) {

  // BITCOIN
  // -------

  if (txHeight == 91812 && txId.String() == "d5d27987d2a3dfc724e359870c6644b40e497bdc0589a033220fe15429d88599") ||
     (txHeight == 91722 && txId.String() == "e3bf3d07d4b0375638d5f1db5255fe07ba2c4cb067cd81b84ee974b6585fb468") {
    return BugBitcoin, true
  }

  // NAMECOIN
  // --------

  if (
    // These transactions have name outputs but a non-Namecoin tx version.
    // They contain NAME_NEWs, which are fine, and also NAME_FIRSTUPDATE.  The
    // latter are not interpreted by namecoind, thus also ignore them for us
    // here. *Just* NAME_NEW outputs are handled by special "lenient version
    // checks". Below are only those transactions that also have name
    // registrations.
    (txHeight == 98423 && txId.String() == "bff3ed6873e5698b97bf0c28c29302b59588590b747787c7d1ef32decdabe0d1") ||
    (txHeight == 98424 && txId.String() == "e9b211007e5cac471769212ca0f47bb066b81966a8e541d44acf0f8a1bd24976") ||
    (txHeight == 98425 && txId.String() == "8aa2b0fc7d1033de28e0192526765a72e9df0c635f7305bdc57cb451ed01a4ca")) {
    return BugFullyIgnore, true
  }

  if (
    // These transaction has both a NAME_NEW and a NAME_FIRSTUPDATE as inputs.
    // This was accepted due to the "argument concatenation" bug.
    // It is fine to accept it as valid and just process the NAME_UPDATE
    // output that builds on the NAME_FIRSTUPDATE input. (NAME_NEW has no
    // special side-effect in applying anyway.)
    (txHeight == 99381 && txId.String() == "774d4c446cecfc40b1c02fdc5a13be6d2007233f9d91daefab6b3c2e70042f05")) {
    return BugFullyApply, true
  }

  // libcoin's name stealing bugs.
  if (txHeight == 139872 && txId.String() == "2f034f2499c136a2c5a922ca4be65c1292815c753bbb100a2a26d5ad532c3919") {
    return BugInUTXO, true
  }

  if (txHeight == 139936 && txId.String() == "c3e76d5384139228221cce60250397d1b87adf7366086bc8d6b5e6eee03c55c7") {
    return BugFullyIgnore, true
  }

  return BugNone, false
}
