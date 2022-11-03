// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const OptimismMintableERC721FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":27366,\"contract\":\"contracts/universal/OptimismMintableERC721Factory.sol:OptimismMintableERC721Factory\",\"label\":\"isOptimismMintableERC721\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var OptimismMintableERC721FactoryStorageLayout = new(solc.StorageLayout)

var OptimismMintableERC721FactoryDeployedBin = "0x60806040523480156200001157600080fd5b5060043610620000875760003560e01c8063d97df6521162000062578063d97df652146200011c578063e78cea921462000159578063e95181961462000180578063ee9a31a214620001a757600080fd5b806354fd4d50146200008c5780635572acae14620000ae5780637d1d0c5b14620000e5575b600080fd5b62000096620001cf565b604051620000a5919062000642565b60405180910390f35b620000d4620000bf36600462000688565b60006020819052908152604090205460ff1681565b6040519015158152602001620000a5565b6200010d7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001620000a5565b620001336200012d36600462000788565b6200027a565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001620000a5565b7f000000000000000000000000000000000000000000000000000000000000000062000133565b7f00000000000000000000000000000000000000000000000000000000000000006200010d565b620001337f000000000000000000000000000000000000000000000000000000000000000081565b6060620001fc7f000000000000000000000000000000000000000000000000000000000000000062000460565b620002277f000000000000000000000000000000000000000000000000000000000000000062000460565b620002527f000000000000000000000000000000000000000000000000000000000000000062000460565b604051602001620002669392919062000805565b604051602081830303815290604052905090565b600073ffffffffffffffffffffffffffffffffffffffff84166200034b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4f7074696d69736d4d696e7461626c65455243373231466163746f72793a204c908201527f3120746f6b656e20616464726573732063616e6e6f742062652061646472657360648201527f7328302900000000000000000000000000000000000000000000000000000000608482015260a40160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000868686604051620003a090620005b5565b620003b095949392919062000881565b604051809103906000f080158015620003cd573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff8181166000818152602081815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590513381529394509188169290917fe72783bb8e0ca31286b85278da59684dd814df9762a52f0837f89edd1483b299910160405180910390a3949350505050565b606081600003620004a457505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115620004d45780620004bb8162000911565b9150620004cc9050600a836200097b565b9150620004a8565b60008167ffffffffffffffff811115620004f257620004f2620006a6565b6040519080825280601f01601f1916602001820160405280156200051d576020820181803683370190505b5090505b8415620005ad576200053560018362000992565b915062000544600a86620009ac565b62000551906030620009c3565b60f81b818381518110620005695762000569620009de565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350620005a5600a866200097b565b945062000521565b949350505050565b6131c28062000a0e83390190565b60005b83811015620005e0578181015183820152602001620005c6565b83811115620005f0576000848401525b50505050565b6000815180845262000610816020860160208601620005c3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620006576020830184620005f6565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200068357600080fd5b919050565b6000602082840312156200069b57600080fd5b62000657826200065e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112620006e757600080fd5b813567ffffffffffffffff80821115620007055762000705620006a6565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156200074e576200074e620006a6565b816040528381528660208588010111156200076857600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200079e57600080fd5b620007a9846200065e565b9250602084013567ffffffffffffffff80821115620007c757600080fd5b620007d587838801620006d5565b93506040860135915080821115620007ec57600080fd5b50620007fb86828701620006d5565b9150509250925092565b6000845162000819818460208901620005c3565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855162000857816001850160208a01620005c3565b6001920191820152835162000874816002840160208801620005c3565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015280861660408401525060a06060830152620008c260a0830185620005f6565b8281036080840152620008d68185620005f6565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620009455762000945620008e2565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826200098d576200098d6200094c565b500490565b600082821015620009a757620009a7620008e2565b500390565b600082620009be57620009be6200094c565b500690565b60008219821115620009d957620009d9620008e2565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe60e06040523480156200001157600080fd5b50604051620031c2380380620031c283398101604081905262000034916200062d565b8181600062000044838262000756565b50600162000053828262000756565b5050506001600160a01b038516620000d85760405162461bcd60e51b815260206004820152603360248201527f4f7074696d69736d4d696e7461626c654552433732313a20627269646765206360448201527f616e6e6f7420626520616464726573732830290000000000000000000000000060648201526084015b60405180910390fd5b83600003620001505760405162461bcd60e51b815260206004820152603660248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465206360448201527f6861696e2069642063616e6e6f74206265207a65726f000000000000000000006064820152608401620000cf565b6001600160a01b038316620001ce5760405162461bcd60e51b815260206004820152603960248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465207460448201527f6f6b656e2063616e6e6f742062652061646472657373283029000000000000006064820152608401620000cf565b60808490526001600160a01b0383811660a081905290861660c0526200020290601462000256602090811b62000ef817901c565b62000218856200041660201b6200113b1760201c565b6040516020016200022b92919062000822565b604051602081830303815290604052600a90816200024a919062000756565b50505050505062000993565b6060600062000267836002620008ac565b62000274906002620008ce565b6001600160401b038111156200028e576200028e62000553565b6040519080825280601f01601f191660200182016040528015620002b9576020820181803683370190505b509050600360fc1b81600081518110620002d757620002d7620008e9565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110620003095762000309620008e9565b60200101906001600160f81b031916908160001a90535060006200032f846002620008ac565b6200033c906001620008ce565b90505b6001811115620003be576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110620003745762000374620008e9565b1a60f81b8282815181106200038d576200038d620008e9565b60200101906001600160f81b031916908160001a90535060049490941c93620003b681620008ff565b90506200033f565b5083156200040f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401620000cf565b9392505050565b6060816000036200043e5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156200046e5780620004558162000919565b9150620004669050600a836200094b565b915062000442565b6000816001600160401b038111156200048b576200048b62000553565b6040519080825280601f01601f191660200182016040528015620004b6576020820181803683370190505b5090505b84156200052e57620004ce60018362000962565b9150620004dd600a866200097c565b620004ea906030620008ce565b60f81b818381518110620005025762000502620008e9565b60200101906001600160f81b031916908160001a90535062000526600a866200094b565b9450620004ba565b949350505050565b80516001600160a01b03811681146200054e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620005865781810151838201526020016200056c565b8381111562000596576000848401525b50505050565b600082601f830112620005ae57600080fd5b81516001600160401b0380821115620005cb57620005cb62000553565b604051601f8301601f19908116603f01168101908282118183101715620005f657620005f662000553565b816040528381528660208588010111156200061057600080fd5b6200062384602083016020890162000569565b9695505050505050565b600080600080600060a086880312156200064657600080fd5b620006518662000536565b945060208601519350620006686040870162000536565b60608701519093506001600160401b03808211156200068657600080fd5b6200069489838a016200059c565b93506080880151915080821115620006ab57600080fd5b50620006ba888289016200059c565b9150509295509295909350565b600181811c90821680620006dc57607f821691505b602082108103620006fd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200075157600081815260208120601f850160051c810160208610156200072c5750805b601f850160051c820191505b818110156200074d5782815560010162000738565b5050505b505050565b81516001600160401b0381111562000772576200077262000553565b6200078a81620007838454620006c7565b8462000703565b602080601f831160018114620007c25760008415620007a95750858301515b600019600386901b1c1916600185901b1785556200074d565b600085815260208120601f198616915b82811015620007f357888601518255948401946001909101908401620007d2565b5085821015620008125787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6832ba3432b932bab69d60b91b8152600083516200084881600985016020880162000569565b600160fe1b60099184019182015283516200086b81600a84016020880162000569565b712f746f6b656e5552493f75696e743235363d60701b600a9290910191820152601c01949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615620008c957620008c962000896565b500290565b60008219821115620008e457620008e462000896565b500190565b634e487b7160e01b600052603260045260246000fd5b60008162000911576200091162000896565b506000190190565b6000600182016200092e576200092e62000896565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826200095d576200095d62000935565b500490565b60008282101562000977576200097762000896565b500390565b6000826200098e576200098e62000935565b500690565b60805160a05160c0516127dc620009e66000396000818161039b0152818161043301528181610b2b0152610c4d0152600081816101d501526103750152600081816102e201526103c101526127dc6000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80637d1d0c5b116100ee578063c87b56dd11610097578063e78cea9211610071578063e78cea9214610399578063e9518196146103bf578063e985e9c5146103e5578063ee9a31a21461042e57600080fd5b8063c87b56dd14610358578063d547cfb71461036b578063d6c0b2c41461037357600080fd5b8063a1448194116100c8578063a14481941461031f578063a22cb46514610332578063b88d4fde1461034557600080fd5b80637d1d0c5b146102dd57806395d89b41146103045780639dc29fac1461030c57600080fd5b806323b872dd116101505780634f6ccce71161012a5780634f6ccce7146102a45780636352211e146102b757806370a08231146102ca57600080fd5b806323b872dd1461026b5780632f745c591461027e57806342842e0e1461029157600080fd5b8063081812fc11610181578063081812fc14610231578063095ea7b31461024457806318160ddd1461025957600080fd5b806301ffc9a7146101a8578063033964be146101d057806306fdde031461021c575b600080fd5b6101bb6101b6366004612229565b610455565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c7565b610224610504565b6040516101c791906122bc565b6101f761023f3660046122cf565b610596565b610257610252366004612311565b6105ca565b005b6008545b6040519081526020016101c7565b61025761027936600461233b565b61075b565b61025d61028c366004612311565b6107fc565b61025761029f36600461233b565b6108cb565b61025d6102b23660046122cf565b6108e6565b6101f76102c53660046122cf565b6109a4565b61025d6102d8366004612377565b610a36565b61025d7f000000000000000000000000000000000000000000000000000000000000000081565b610224610b04565b61025761031a366004612311565b610b13565b61025761032d366004612311565b610c35565b610257610340366004612392565b610d4c565b6102576103533660046123fd565b610d5b565b6102246103663660046122cf565b610e03565b610224610e6a565b7f00000000000000000000000000000000000000000000000000000000000000006101f7565b7f00000000000000000000000000000000000000000000000000000000000000006101f7565b7f000000000000000000000000000000000000000000000000000000000000000061025d565b6101bb6103f33660046124f7565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f74259ebf000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000084168214806104ed57507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b806104fc57506104fc84611270565b949350505050565b6060600080546105139061252a565b80601f016020809104026020016040519081016040528092919081815260200182805461053f9061252a565b801561058c5780601f106105615761010080835404028352916020019161058c565b820191906000526020600020905b81548152906001019060200180831161056f57829003601f168201915b5050505050905090565b60006105a1826112c6565b5060009081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60006105d5826109a4565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610697576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff821614806106c057506106c081336103f3565b61074c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000606482015260840161068e565b6107568383611354565b505050565b61076533826113f4565b6107f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f766564000000000000000000000000000000000000606482015260840161068e565b6107568383836114b3565b600061080783610a36565b8210610895576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e6473000000000000000000000000000000000000000000606482015260840161068e565b5073ffffffffffffffffffffffffffffffffffffffff919091166000908152600660209081526040808320938352929052205490565b61075683838360405180602001604052806000815250610d5b565b60006108f160085490565b821061097f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e64730000000000000000000000000000000000000000606482015260840161068e565b600882815481106109925761099261257d565b90600052602060002001549050919050565b60008181526002602052604081205473ffffffffffffffffffffffffffffffffffffffff1680610a30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e2049440000000000000000604482015260640161068e565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff8216610adb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e65720000000000000000000000000000000000000000000000606482015260840161068e565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6060600180546105139061252a565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e000000000000606482015260840161068e565b610be181611725565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610c2991815260200190565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610cfa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e000000000000606482015260840161068e565b610d0482826117fe565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688582604051610c2991815260200190565b610d57338383611818565b5050565b610d6533836113f4565b610df1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f766564000000000000000000000000000000000000606482015260840161068e565b610dfd84848484611945565b50505050565b6060610e0e826112c6565b6000610e186119e8565b90506000815111610e385760405180602001604052806000815250610e63565b80610e428461113b565b604051602001610e539291906125ac565b6040516020818303038152906040525b9392505050565b600a8054610e779061252a565b80601f0160208091040260200160405190810160405280929190818152602001828054610ea39061252a565b8015610ef05780601f10610ec557610100808354040283529160200191610ef0565b820191906000526020600020905b815481529060010190602001808311610ed357829003601f168201915b505050505081565b60606000610f0783600261260a565b610f12906002612647565b67ffffffffffffffff811115610f2a57610f2a6123ce565b6040519080825280601f01601f191660200182016040528015610f54576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610f8b57610f8b61257d565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610fee57610fee61257d565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600061102a84600261260a565b611035906001612647565b90505b60018111156110d2577f303132333435363738396162636465660000000000000000000000000000000085600f16601081106110765761107661257d565b1a60f81b82828151811061108c5761108c61257d565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c936110cb8161265f565b9050611038565b508315610e63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161068e565b60608160000361117e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156111a8578061119281612694565b91506111a19050600a836126fb565b9150611182565b60008167ffffffffffffffff8111156111c3576111c36123ce565b6040519080825280601f01601f1916602001820160405280156111ed576020820181803683370190505b5090505b84156104fc5761120260018361270f565b915061120f600a86612726565b61121a906030612647565b60f81b81838151811061122f5761122f61257d565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611269600a866126fb565b94506111f1565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f780e9d63000000000000000000000000000000000000000000000000000000001480610a305750610a30826119f7565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16611351576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e2049440000000000000000604482015260640161068e565b50565b600081815260046020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841690811790915581906113ae826109a4565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080611400836109a4565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061146e575073ffffffffffffffffffffffffffffffffffffffff80821660009081526005602090815260408083209388168352929052205460ff165b806104fc57508373ffffffffffffffffffffffffffffffffffffffff1661149484610596565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b8273ffffffffffffffffffffffffffffffffffffffff166114d3826109a4565b73ffffffffffffffffffffffffffffffffffffffff1614611576576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e6572000000000000000000000000000000000000000000000000000000606482015260840161068e565b73ffffffffffffffffffffffffffffffffffffffff8216611618576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161068e565b611623838383611ada565b61162e600082611354565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080546001929061166490849061270f565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080546001929061169f908490612647565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6000611730826109a4565b905061173e81600084611ada565b611749600083611354565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260036020526040812080546001929061177f90849061270f565b909155505060008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555183919073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b610d57828260405180602001604052806000815250611be0565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036118ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161068e565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6119508484846114b3565b61195c84848484611c83565b610dfd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e7465720000000000000000000000000000606482015260840161068e565b6060600a80546105139061252a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd000000000000000000000000000000000000000000000000000000001480611a8a57507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b80610a3057507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610a30565b73ffffffffffffffffffffffffffffffffffffffff8316611b4257611b3d81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b611b7f565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614611b7f57611b7f8382611e76565b73ffffffffffffffffffffffffffffffffffffffff8216611ba35761075681611f2d565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610756576107568282611fdc565b611bea838361202d565b611bf76000848484611c83565b610756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e7465720000000000000000000000000000606482015260840161068e565b600073ffffffffffffffffffffffffffffffffffffffff84163b15611e6b576040517f150b7a0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290611cfa90339089908890889060040161273a565b6020604051808303816000875af1925050508015611d53575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611d5091810190612783565b60015b611e20573d808015611d81576040519150601f19603f3d011682016040523d82523d6000602084013e611d86565b606091505b508051600003611e18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e7465720000000000000000000000000000606482015260840161068e565b805181602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000167f150b7a02000000000000000000000000000000000000000000000000000000001490506104fc565b506001949350505050565b60006001611e8384610a36565b611e8d919061270f565b600083815260076020526040902054909150808214611eed5773ffffffffffffffffffffffffffffffffffffffff841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b50600091825260076020908152604080842084905573ffffffffffffffffffffffffffffffffffffffff9094168352600681528383209183525290812055565b600854600090611f3f9060019061270f565b60008381526009602052604081205460088054939450909284908110611f6757611f6761257d565b906000526020600020015490508060088381548110611f8857611f8861257d565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480611fc057611fc06127a0565b6001900381819060005260206000200160009055905550505050565b6000611fe783610a36565b73ffffffffffffffffffffffffffffffffffffffff9093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b73ffffffffffffffffffffffffffffffffffffffff82166120aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161068e565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1615612136576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161068e565b61214260008383611ada565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805460019290612178908490612647565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461135157600080fd5b60006020828403121561223b57600080fd5b8135610e63816121fb565b60005b83811015612261578181015183820152602001612249565b83811115610dfd5750506000910152565b6000815180845261228a816020860160208601612246565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e636020830184612272565b6000602082840312156122e157600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461230c57600080fd5b919050565b6000806040838503121561232457600080fd5b61232d836122e8565b946020939093013593505050565b60008060006060848603121561235057600080fd5b612359846122e8565b9250612367602085016122e8565b9150604084013590509250925092565b60006020828403121561238957600080fd5b610e63826122e8565b600080604083850312156123a557600080fd5b6123ae836122e8565b9150602083013580151581146123c357600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000806080858703121561241357600080fd5b61241c856122e8565b935061242a602086016122e8565b925060408501359150606085013567ffffffffffffffff8082111561244e57600080fd5b818701915087601f83011261246257600080fd5b813581811115612474576124746123ce565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156124ba576124ba6123ce565b816040528281528a60208487010111156124d357600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561250a57600080fd5b612513836122e8565b9150612521602084016122e8565b90509250929050565b600181811c9082168061253e57607f821691505b602082108103612577577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600083516125be818460208801612246565b8351908301906125d2818360208801612246565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612642576126426125db565b500290565b6000821982111561265a5761265a6125db565b500190565b60008161266e5761266e6125db565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036126c5576126c56125db565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261270a5761270a6126cc565b500490565b600082821015612721576127216125db565b500390565b600082612735576127356126cc565b500690565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526127796080830184612272565b9695505050505050565b60006020828403121561279557600080fd5b8151610e63816121fb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC721FactoryStorageLayoutJSON), OptimismMintableERC721FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryStorageLayout
	deployedBytecodes["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryDeployedBin
}
